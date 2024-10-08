package instance

import (
	"bytes"
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/dustin/go-humanize"
	"github.com/scaleway/scaleway-cli/v2/internal/core"
	block "github.com/scaleway/scaleway-sdk-go/api/block/v1alpha1"
	"github.com/scaleway/scaleway-sdk-go/api/instance/v1"
	"github.com/scaleway/scaleway-sdk-go/api/marketplace/v2"
	"github.com/scaleway/scaleway-sdk-go/logger"
	"github.com/scaleway/scaleway-sdk-go/scw"
	"github.com/scaleway/scaleway-sdk-go/validation"
)

type instanceCreateServerRequest struct {
	Zone              scw.Zone
	ProjectID         *string
	Image             string
	Type              string
	Name              string
	RootVolume        string
	AdditionalVolumes []string
	IP                string
	Tags              []string
	IPv6              bool
	Stopped           bool
	SecurityGroupID   string
	PlacementGroupID  string

	// Windows
	AdminPasswordEncryptionSSHKeyID *string

	// IP Mobility
	RoutedIPEnabled *bool

	// Deprecated
	BootscriptID string
	CloudInit    string
	BootType     string

	// Deprecated, use project-id instead
	OrganizationID *string
}

func serverCreateCommand() *core.Command {
	return &core.Command{
		Short:     `Create server`,
		Long:      `Create an instance server.`,
		Namespace: "instance",
		Verb:      "create",
		Resource:  "server",
		ArgsType:  reflect.TypeOf(instanceCreateServerRequest{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:             "image",
				Short:            "Image ID or label of the server",
				Default:          core.DefaultValueSetter("ubuntu_jammy"),
				Required:         true,
				AutoCompleteFunc: instanceServerCreateImageAutoCompleteFunc,
			},
			{
				Name:     "type",
				Short:    "Server commercial type (help: https://www.scaleway.com/en/docs/compute/instances/reference-content/choosing-instance-type/)",
				Default:  core.DefaultValueSetter("DEV1-S"),
				Required: true,
				ValidateFunc: func(_ *core.ArgSpec, _ interface{}) error {
					// Allow all commercial types
					return nil
				},
				AutoCompleteFunc: completeServerType,
			},
			{
				Name:    "name",
				Short:   "Server name",
				Default: core.RandomValueGenerator("srv"),
			},
			{
				Name:  "root-volume",
				Short: "Local root volume of the server",
			},
			{
				Name:  "additional-volumes.{index}",
				Short: "Additional local and block volumes attached to your server",
			},
			{
				Name:    "ip",
				Short:   `Either an IP, an IP ID, 'new' to create a new IP, 'dynamic' to use a dynamic IP or 'none' for no public IP (new | dynamic | none | <id> | <address>)`,
				Default: core.DefaultValueSetter("new"),
			},
			{
				Name:  "tags.{index}",
				Short: "Server tags",
			},
			{
				Name:  "ipv6",
				Short: "Enable IPv6, to be used with routed-ip-enabled=false",
			},
			{
				Name:  "stopped",
				Short: "Do not start server after its creation",
			},
			{
				Name:  "security-group-id",
				Short: "The security group ID used for this server",
			},
			{
				Name:  "placement-group-id",
				Short: "The placement group ID in which the server has to be created",
			},
			{
				Name:        "cloud-init",
				Short:       "The cloud-init script to use",
				CanLoadFile: true,
			},
			{
				Name:       "boot-type",
				Short:      "The boot type to use, if empty the local boot will be used. Will be overwritten to bootscript if bootscript-id is set.",
				Default:    core.DefaultValueSetter(instance.BootTypeLocal.String()),
				EnumValues: []string{instance.BootTypeLocal.String(), instance.BootTypeBootscript.String(), instance.BootTypeRescue.String()},
			},
			{
				Name:  "routed-ip-enabled",
				Short: "Enable routed IP support",
			},
			{
				Name:             "admin-password-encryption-ssh-key-id",
				Short:            "ID of the IAM SSH Key used to encrypt generated admin password. Required when creating a windows server.",
				AutoCompleteFunc: completeSSHKeyID,
			},
			core.ProjectIDArgSpec(),
			core.ZoneArgSpec((*instance.API)(nil).Zones()...),
			core.OrganizationIDArgSpec(),
		},
		Run:      instanceServerCreateRun,
		WaitFunc: instanceWaitServerCreateRun(),
		SeeAlsos: []*core.SeeAlso{{
			Short:   "List marketplace label images",
			Command: "scw marketplace image list",
		}},
		Examples: []*core.Example{
			{
				Short:    "Create and start an instance on Ubuntu Focal",
				ArgsJSON: `{"image":"ubuntu_focal","start":true}`,
			},
			{
				Short:    "Create a GP1-XS instance, give it a name and add tags",
				ArgsJSON: `{"image":"ubuntu_focal","type":"GP1-XS","name":"foo","tags":["prod","blue"]}`,
			},
			{
				Short:    "Create an instance with 2 additional block volumes (50GB and 100GB)",
				ArgsJSON: `{"image":"ubuntu_focal","additional_volumes":["block:50GB","block:100GB"]}`,
			},
			{
				Short:    "Create an instance with 2 local volumes (10GB and 10GB)",
				ArgsJSON: `{"image":"ubuntu_focal","root_volume":"local:10GB","additional_volumes":["local:10GB"]}`,
			},
			{
				Short:    "Create an instance with volumes from snapshots",
				ArgsJSON: `{"image":"ubuntu_focal","root_volume":"local:<snapshot_id>","additional_volumes":["block:<snapshot_id>"]}`,
			},
			{
				Short:    "Create and start an instance from a snapshot",
				ArgsJSON: `{"image":"none","root_volume":"local:<snapshot_id>"}`,
			},
			{
				Short:    "Create and start an instance using existing volume",
				ArgsJSON: `{"image":"ubuntu_focal","additional_volumes":["<volume_id>"]}`,
			},
			{
				Short: "Use an existing IP",
				Raw: `ip=$(scw instance ip create | grep id | awk '{ print $2 }')
scw instance server create image=ubuntu_focal ip=$ip`,
			},
		},
	}
}

func instanceWaitServerCreateRun() core.WaitFunc {
	return func(ctx context.Context, argsI, respI interface{}) (interface{}, error) {
		return instance.NewAPI(core.ExtractClient(ctx)).WaitForServer(&instance.WaitForServerRequest{
			Zone:          argsI.(*instanceCreateServerRequest).Zone,
			ServerID:      respI.(*instance.Server).ID,
			Timeout:       scw.TimeDurationPtr(serverActionTimeout),
			RetryInterval: core.DefaultRetryInterval,
		})
	}
}

func instanceServerCreateRun(ctx context.Context, argsI interface{}) (i interface{}, e error) {
	var err error
	args := argsI.(*instanceCreateServerRequest)

	//
	// STEP 1: Argument handling and API requests creation.
	//

	client := core.ExtractClient(ctx)

	serverBuilder := NewServerBuilder(client, args.Name, args.Zone, args.Type).
		AddOrganizationID(args.OrganizationID).
		AddProjectID(args.ProjectID).
		AddEnableIPv6(scw.BoolPtr(args.IPv6)).
		AddTags(args.Tags).
		AddRoutedIPEnabled(args.RoutedIPEnabled).
		AddAdminPasswordEncryptionSSHKeyID(args.AdminPasswordEncryptionSSHKeyID).
		AddBootType(args.BootType).
		AddSecurityGroup(args.SecurityGroupID).
		AddPlacementGroup(args.PlacementGroupID)

	serverBuilder, err = serverBuilder.AddVolumes(args.RootVolume, args.AdditionalVolumes)
	if err != nil {
		return nil, err
	}

	serverBuilder, err = serverBuilder.AddImage(args.Image)
	if err != nil {
		return nil, err
	}

	serverBuilder, err = serverBuilder.AddIP(args.IP)
	if err != nil {
		return nil, err
	}

	//
	// STEP 2: Validation and requests
	//

	err = serverBuilder.Validate()
	if err != nil {
		return nil, err
	}

	createReq, createIPReq := serverBuilder.Build()
	needIPCreation := createIPReq != nil

	//
	// IP creation
	//
	apiInstance := instance.NewAPI(client)

	if needIPCreation {
		logger.Debugf("creating IP")

		ipRes, err := apiInstance.CreateIP(createIPReq)
		if err != nil {
			return nil, fmt.Errorf("error while creating your public IP: %s", err)
		}
		createReq.PublicIP = scw.StringPtr(ipRes.IP.ID)
		logger.Debugf("IP created: %s", createReq.PublicIP)
	}

	//
	// Server Creation
	//
	logger.Debugf("creating server")
	serverRes, err := apiInstance.CreateServer(createReq)
	if err != nil {
		if needIPCreation && createReq.PublicIP != nil {
			// Delete the created IP
			logger.Debugf("deleting created IP: %s", createReq.PublicIP)
			err := apiInstance.DeleteIP(&instance.DeleteIPRequest{
				Zone: args.Zone,
				IP:   *createReq.PublicIP,
			})
			if err != nil {
				logger.Warningf("cannot delete the create IP %s: %s.", createReq.PublicIP, err)
			}
		}

		return nil, fmt.Errorf("cannot create the server: %s", err)
	}
	server := serverRes.Server
	logger.Debugf("server created %s", server.ID)

	//
	// Cloud-init
	//
	if args.CloudInit != "" {
		err := apiInstance.SetServerUserData(&instance.SetServerUserDataRequest{
			Zone:     args.Zone,
			ServerID: server.ID,
			Key:      "cloud-init",
			Content:  bytes.NewBufferString(args.CloudInit),
		})
		if err != nil {
			logger.Warningf("error while setting up your cloud-init metadata: %s. Note that the server is successfully created.", err)
		} else {
			logger.Debugf("cloud-init set")
		}
	}

	//
	// Start server by default
	//
	if !args.Stopped {
		logger.Debugf("starting server")
		_, err := apiInstance.ServerAction(&instance.ServerActionRequest{
			Zone:     args.Zone,
			ServerID: server.ID,
			Action:   instance.ServerActionPoweron,
		})
		if err != nil {
			logger.Warningf("Cannot start the server: %s. Note that the server is successfully created.", err)
		} else {
			logger.Debugf("server started")
		}
	}

	return server, nil
}

func addDefaultVolumes(serverType *instance.ServerType, volumes map[string]*instance.VolumeServerTemplate) map[string]*instance.VolumeServerTemplate {
	needScratch := false
	hasScratch := false
	defaultVolumes := []*instance.VolumeServerTemplate(nil)
	if serverType.ScratchStorageMaxSize != nil && *serverType.ScratchStorageMaxSize > 0 {
		needScratch = true
	}
	for _, volume := range volumes {
		if volume.VolumeType == instance.VolumeVolumeTypeScratch {
			hasScratch = true
		}
	}

	if needScratch && !hasScratch {
		defaultVolumes = append(defaultVolumes, &instance.VolumeServerTemplate{
			Name:       scw.StringPtr("default-cli-scratch-volume"),
			Size:       serverType.ScratchStorageMaxSize,
			VolumeType: instance.VolumeVolumeTypeScratch,
		})
	}

	if defaultVolumes != nil {
		if volumes == nil {
			volumes = make(map[string]*instance.VolumeServerTemplate)
		}
		maxKey := 1
		for k := range volumes {
			key, err := strconv.Atoi(k)
			if err == nil && key > maxKey {
				maxKey = key
			}
		}
		for i, vol := range defaultVolumes {
			volumes[strconv.Itoa(maxKey+i)] = vol
		}
	}

	return volumes
}

// buildVolumes creates the initial volume map.
// It is not the definitive one, it will be mutated all along the process.
func buildVolumes(api *instance.API, blockAPI *block.API, zone scw.Zone, serverName, rootVolume string, additionalVolumes []string) (map[string]*instance.VolumeServerTemplate, error) {
	volumes := make(map[string]*instance.VolumeServerTemplate)
	if rootVolume != "" {
		rootVolumeTemplate, err := buildVolumeTemplate(api, blockAPI, zone, rootVolume)
		if err != nil {
			return nil, err
		}

		volumes["0"] = rootVolumeTemplate
	}

	for i, v := range additionalVolumes {
		volumeTemplate, err := buildVolumeTemplate(api, blockAPI, zone, v)
		if err != nil {
			return nil, err
		}
		index := strconv.Itoa(i + 1)
		volumeTemplate.Name = scw.StringPtr(serverName + "-" + index)

		volumes[index] = volumeTemplate
	}

	return volumes, nil
}

// buildVolumeTemplate creates a instance.VolumeTemplate from a 'volumes' argument item.
//
// Volumes definition must be through multiple arguments (eg: volumes.0="l:20GB" volumes.1="b:100GB")
//
// A valid volume format is either
// - a "creation" format: ^((local|l|block|b|scratch|s):)?\d+GB?$ (size is handled by go-humanize, so other sizes are supported)
// - a "creation" format with a snapshot id: l:<uuid> b:<uuid>
// - a UUID format
func buildVolumeTemplate(api *instance.API, blockAPI *block.API, zone scw.Zone, flagV string) (*instance.VolumeServerTemplate, error) {
	parts := strings.Split(strings.TrimSpace(flagV), ":")

	// Create volume.
	if len(parts) == 2 {
		vt := &instance.VolumeServerTemplate{}

		switch parts[0] {
		case "l", "local":
			vt.VolumeType = instance.VolumeVolumeTypeLSSD
		case "b", "block":
			vt.VolumeType = instance.VolumeVolumeTypeBSSD
		case "s", "scratch":
			vt.VolumeType = instance.VolumeVolumeTypeScratch
		case "sbs":
			vt.VolumeType = instance.VolumeVolumeTypeSbsVolume
		default:
			return nil, fmt.Errorf("invalid volume type %s in %s volume", parts[0], flagV)
		}

		if validation.IsUUID(parts[1]) {
			return buildVolumeTemplateFromSnapshot(api, zone, parts[1], vt.VolumeType)
		}

		size, err := humanize.ParseBytes(parts[1])
		if err != nil {
			return nil, fmt.Errorf("invalid size format %s in %s volume", parts[1], flagV)
		}
		vt.Size = scw.SizePtr(scw.Size(size))

		return vt, nil
	}

	// UUID format.
	if len(parts) == 1 && validation.IsUUID(parts[0]) {
		return buildVolumeTemplateFromUUID(api, blockAPI, zone, parts[0])
	}

	return nil, &core.CliError{
		Err:     fmt.Errorf("invalid volume format '%s'", flagV),
		Details: "",
		Hint:    `You must provide either a UUID ("11111111-1111-1111-1111-111111111111"), a local volume size ("local:100G" or "l:100G") or a block volume size ("block:100G" or "b:100G").`,
	}
}

// buildVolumeTemplateFromUUID validate an UUID volume and add their types and sizes.
// Add volume types and sizes allow US to treat UUID volumes like the others and simplify the implementation.
// The instance API refuse the type and the size for UUID volumes, therefore,
// sanitizeVolumeMap function will remove them.
func buildVolumeTemplateFromUUID(api *instance.API, blockAPI *block.API, zone scw.Zone, volumeUUID string) (*instance.VolumeServerTemplate, error) {
	res, err := api.GetVolume(&instance.GetVolumeRequest{
		Zone:     zone,
		VolumeID: volumeUUID,
	})
	if err != nil && !core.IsNotFoundError(err) {
		return nil, err
	}

	if res != nil {
		// Check that volume is not already attached to a server.
		if res.Volume.Server != nil {
			return nil, fmt.Errorf("volume %s is already attached to %s server", res.Volume.ID, res.Volume.Server.ID)
		}

		return &instance.VolumeServerTemplate{
			ID:         &res.Volume.ID,
			VolumeType: res.Volume.VolumeType,
			Size:       &res.Volume.Size,
		}, nil
	}

	blockRes, err := blockAPI.GetVolume(&block.GetVolumeRequest{
		Zone:     zone,
		VolumeID: volumeUUID,
	})
	if err != nil {
		if core.IsNotFoundError(err) {
			return nil, fmt.Errorf("volume %s does not exist", volumeUUID)
		}
		return nil, err
	}

	if len(blockRes.References) > 0 {
		return nil, fmt.Errorf("volume %s is already attached to %s %s", blockRes.ID, blockRes.References[0].ProductResourceID, blockRes.References[0].ProductResourceType)
	}

	return &instance.VolumeServerTemplate{
		ID:         &blockRes.ID,
		VolumeType: instance.VolumeVolumeTypeSbsVolume, // TODO: support snapshot
	}, nil
}

// buildVolumeTemplateFromUUID validate a snapshot UUID and check that requested volume type is compatible.
// The instance API refuse the size for Snapshot volumes, therefore,
// sanitizeVolumeMap function will remove them.
func buildVolumeTemplateFromSnapshot(api *instance.API, zone scw.Zone, snapshotUUID string, volumeType instance.VolumeVolumeType) (*instance.VolumeServerTemplate, error) {
	res, err := api.GetSnapshot(&instance.GetSnapshotRequest{
		Zone:       zone,
		SnapshotID: snapshotUUID,
	})
	if err != nil {
		if core.IsNotFoundError(err) {
			return nil, fmt.Errorf("snapshot %s does not exist", snapshotUUID)
		}
		return nil, err
	}

	snapshotType := res.Snapshot.VolumeType

	if snapshotType != instance.VolumeVolumeTypeUnified && snapshotType != volumeType {
		return nil, fmt.Errorf("snapshot of type %s not compatible with requested volume type %s", snapshotType, volumeType)
	}

	return &instance.VolumeServerTemplate{
		Name:         &res.Snapshot.Name,
		VolumeType:   volumeType,
		BaseSnapshot: &res.Snapshot.ID,
		Size:         &res.Snapshot.Size,
	}, nil
}

func validateImageServerTypeCompatibility(image *instance.Image, serverType *instance.ServerType, CommercialType string) error {
	// An instance might not have any constraints on the local volume size
	if serverType.VolumesConstraint.MaxSize == 0 {
		return nil
	}
	if image.RootVolume.VolumeType == instance.VolumeVolumeTypeLSSD && image.RootVolume.Size > serverType.VolumesConstraint.MaxSize {
		return fmt.Errorf("image %s requires %s on root volume, but root volume is constrained between %s and %s on %s",
			image.ID,
			humanize.Bytes(uint64(image.RootVolume.Size)),
			humanize.Bytes(uint64(serverType.VolumesConstraint.MinSize)),
			humanize.Bytes(uint64(serverType.VolumesConstraint.MaxSize)),
			CommercialType,
		)
	}

	return nil
}

// validateLocalVolumeSizes validates the total size of local volumes.
func validateLocalVolumeSizes(volumes map[string]*instance.VolumeServerTemplate, serverType *instance.ServerType, commercialType string, defaultRootVolumeSize scw.Size) error {
	// Calculate local volume total size.
	var localVolumeTotalSize scw.Size
	for _, volume := range volumes {
		if volume.VolumeType == instance.VolumeVolumeTypeLSSD && volume.Size != nil {
			localVolumeTotalSize += *volume.Size
		}
	}

	volumeConstraint := serverType.VolumesConstraint

	// If no root volume provided, count the default root volume size added by the API.
	// Only count if server allows LSSD.
	if rootVolume := volumes["0"]; rootVolume == nil &&
		serverType.PerVolumeConstraint != nil &&
		serverType.PerVolumeConstraint.LSSD != nil &&
		serverType.PerVolumeConstraint.LSSD.MaxSize > 0 {
		localVolumeTotalSize += defaultRootVolumeSize // defaultRootVolumeSize may be used for a block volume
	}

	if localVolumeTotalSize < volumeConstraint.MinSize || localVolumeTotalSize > volumeConstraint.MaxSize {
		minSize := humanize.Bytes(uint64(volumeConstraint.MinSize))
		computedLocal := humanize.Bytes(uint64(localVolumeTotalSize))
		if volumeConstraint.MinSize == volumeConstraint.MaxSize {
			return fmt.Errorf("%s total local volume size must be equal to %s, got %s", commercialType, minSize, computedLocal)
		}

		maxSize := humanize.Bytes(uint64(volumeConstraint.MaxSize))
		return fmt.Errorf("%s total local volume size must be between %s and %s, got %s", commercialType, minSize, maxSize, computedLocal)
	}

	return nil
}

func validateRootVolume(imageRequiredSize scw.Size, rootVolume *instance.VolumeServerTemplate) error {
	if rootVolume == nil {
		return nil
	}

	if rootVolume.ID != nil {
		return &core.CliError{
			Err:     fmt.Errorf("you cannot use an existing volume as a root volume"),
			Details: "You must create an image of this volume and use its ID in the 'image' argument.",
		}
	}

	if rootVolume.Size != nil && *rootVolume.Size < imageRequiredSize {
		return fmt.Errorf("first volume size must be at least %s for this image", humanize.Bytes(uint64(imageRequiredSize)))
	}

	return nil
}

// sanitizeVolumeMap removes extra data for API validation.
func sanitizeVolumeMap(serverName string, volumes map[string]*instance.VolumeServerTemplate) map[string]*instance.VolumeServerTemplate {
	m := make(map[string]*instance.VolumeServerTemplate)

	for index, v := range volumes {
		v.Name = scw.StringPtr(serverName + "-" + index)

		// Remove extra data for API validation.
		switch {
		case v.ID != nil:
			if v.VolumeType == instance.VolumeVolumeTypeSbsVolume {
				v = &instance.VolumeServerTemplate{
					ID:         v.ID,
					VolumeType: v.VolumeType,
				}
			} else {
				v = &instance.VolumeServerTemplate{
					ID:   v.ID,
					Name: v.Name,
				}
			}
		case v.BaseSnapshot != nil:
			v = &instance.VolumeServerTemplate{
				BaseSnapshot: v.BaseSnapshot,
				Name:         v.Name,
				VolumeType:   v.VolumeType,
			}
		case index == "0" && v.Size != nil:
			v = &instance.VolumeServerTemplate{
				VolumeType: v.VolumeType,
				Size:       v.Size,
			}
		}
		m[index] = v
	}

	return m
}

// Caching listImage response for shell completion
var completeListImagesCache *marketplace.ListImagesResponse

func instanceServerCreateImageAutoCompleteFunc(ctx context.Context, prefix string, _ any) core.AutocompleteSuggestions {
	suggestions := core.AutocompleteSuggestions(nil)

	client := core.ExtractClient(ctx)
	api := marketplace.NewAPI(client)

	if completeListImagesCache == nil {
		res, err := api.ListImages(&marketplace.ListImagesRequest{}, scw.WithAllPages())
		if err != nil {
			return nil
		}
		completeListImagesCache = res
	}

	prefix = strings.ToLower(strings.Replace(prefix, "-", "_", -1))

	for _, image := range completeListImagesCache.Images {
		if strings.HasPrefix(image.Label, prefix) {
			suggestions = append(suggestions, image.Label)
		}
	}

	return suggestions
}

// getServerType is a util to get a instance.ServerType by its commercialType
func getServerType(apiInstance *instance.API, zone scw.Zone, commercialType string) *instance.ServerType {
	serverType := (*instance.ServerType)(nil)

	serverTypesRes, err := apiInstance.ListServersTypes(&instance.ListServersTypesRequest{
		Zone: zone,
	}, scw.WithAllPages())
	if err != nil {
		logger.Warningf("cannot get server types: %s", err)
	} else {
		serverType = serverTypesRes.Servers[commercialType]
		if serverType == nil {
			logger.Warningf("unrecognized server type: %s", commercialType)
		}
	}

	return serverType
}
