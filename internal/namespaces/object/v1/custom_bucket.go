package object

import (
	"context"
	"fmt"
	"reflect"

	"github.com/minio/minio-go/v7"
	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

type objectBucketBasicArgs struct {
	Region scw.Region
	Name   string
}

func bucketCreateCommand() *core.Command {
	return &core.Command{
		Namespace: "object",
		Resource:  "bucket",
		Verb:      "create",
		Short:     "Creates a S3 bucket",
		Long:      "Creates an Object Storage Bucket with the S3 protocol. Its name must be unique xxxxxxx",
		ArgsType:  reflect.TypeOf(objectBucketBasicArgs{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Positional: true,
				Required:   true,
				Short:      "The unique name of the bucket",
			},
			core.RegionArgSpec(),
		},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			args := argsI.(*objectBucketBasicArgs)

			minioClient := newMinioClient(ctx, args.Region)
			if minioClient == nil {
				return nil, fmt.Errorf("failed to create client")
			}
			err := minioClient.MakeBucket(ctx, args.Name, minio.MakeBucketOptions{
				Region:        args.Region.String(),
				ObjectLocking: false,
			})

			//s3client := newS3Client(ctx, args.Region)
			//bucket, err := s3client.CreateBucket(ctx, &s3.CreateBucketInput{
			//	Bucket: &args.Name,
			//	//ACL:                        "",
			//	CreateBucketConfiguration: &types.CreateBucketConfiguration{
			//		LocationConstraint: types.BucketLocationConstraint(args.Region),
			//	},
			//	ObjectLockEnabledForBucket: scw.BoolPtr(true),
			//})
			if err != nil {
				return nil, err
			}

			return &core.SuccessResult{
				Message:  "Bucket successfully created",
				Details:  "",
				Resource: "",
				Verb:     "",
				Empty:    false,
				//TargetResource: bucket,
			}, nil
		},
	}
}

func bucketDeleteCommand() *core.Command {
	return &core.Command{
		Namespace: "object",
		Resource:  "bucket",
		Verb:      "delete",
		Short:     "Deletes a S3 bucket",
		Long:      "Deletes ",
		ArgsType:  reflect.TypeOf(objectBucketBasicArgs{}),
		ArgSpecs: core.ArgSpecs{
			{
				Name:       "name",
				Positional: true,
				Required:   true,
				Short:      "The unique name of the bucket",
			},
			core.RegionArgSpec(),
		},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			args := argsI.(*objectBucketBasicArgs)

			minioClient := newMinioClient(ctx, args.Region)
			if minioClient == nil {
				return nil, fmt.Errorf("failed to create client")
			}
			err := minioClient.RemoveBucket(ctx, args.Name)

			//s3client := newS3Client(ctx, args.Region)
			//_, err := s3client.DeleteBucket(ctx, &s3.DeleteBucketInput{
			//	Bucket: &args.Name,
			//	//ExpectedBucketOwner: nil,
			//})
			if err != nil {
				return nil, err
			}

			return nil, nil
		},
		//WaitFunc: func(ctx context.Context, argsI, respI interface{}) (interface{}, error) {
		//	return nil, nil
		//},
	}
}

func bucketGetCommand() *core.Command {
	return &core.Command{
		Namespace: "object",
		Resource:  "bucket",
		Verb:      "get",
		Short:     "",
		Long:      "",
		//ArgsType:  reflect.TypeOf(),
		ArgSpecs: core.ArgSpecs{},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			return nil, nil
		},
		WaitFunc: func(ctx context.Context, argsI, respI interface{}) (interface{}, error) {
			return nil, nil
		},
	}
}

func bucketUpdateCommand() *core.Command {
	return &core.Command{
		Namespace: "object",
		Resource:  "bucket",
		Verb:      "update",
		Short:     "",
		Long:      "",
		//ArgsType:  reflect.TypeOf(),
		ArgSpecs: core.ArgSpecs{},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			return nil, nil
		},
		WaitFunc: func(ctx context.Context, argsI, respI interface{}) (interface{}, error) {
			return nil, nil
		},
	}
}

func bucketEnableVersioningCommand() *core.Command {
	return &core.Command{
		Namespace: "object",
		Resource:  "bucket",
		Verb:      "enable-versioning",
		Short:     "",
		Long:      "",
		//ArgsType:  reflect.TypeOf(),
		ArgSpecs: core.ArgSpecs{},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			return nil, nil
		},
		WaitFunc: func(ctx context.Context, argsI, respI interface{}) (interface{}, error) {
			return nil, nil
		},
	}
}

func bucketDisableVersioningCommand() *core.Command {
	return &core.Command{
		Namespace: "object",
		Resource:  "bucket",
		Verb:      "disable-versioning",
		Short:     "",
		Long:      "",
		//ArgsType:  reflect.TypeOf(),
		ArgSpecs: core.ArgSpecs{},
		Run: func(ctx context.Context, argsI interface{}) (interface{}, error) {
			return nil, nil
		},
		WaitFunc: func(ctx context.Context, argsI, respI interface{}) (interface{}, error) {
			return nil, nil
		},
	}
}
