package object

import (
	"context"
	"os"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

func newMinioClient(ctx context.Context, region scw.Region) *minio.Client {
	scwClient := core.ExtractClient(ctx)
	accessKey, ok := scwClient.GetAccessKey()
	if !ok {
		return nil
	}
	secretKey, ok := scwClient.GetSecretKey()
	if !ok {
		return nil
	}

	customEndpoint := ""
	if ep := os.Getenv("SCW_S3_ENDPOINT"); ep != "" {
		customEndpoint = ep
	} else {
		customEndpoint = "s3." + region.String() + ".scw.cloud"
	}

	minioClient, err := minio.New(customEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: true,
		//Transport:          nil,
		//Trace:              nil,
		//Region:             "",
		//BucketLookup:       0,
		//CustomRegionViaURL: nil,
		//TrailingHeaders:    false,
		//CustomMD5:          nil,
		//CustomSHA256:       nil,
	})
	if err != nil {
		return nil
	}

	return minioClient
}
