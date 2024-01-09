package object

import (
	"context"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/scaleway/scaleway-cli/v2/internal/core"
	"github.com/scaleway/scaleway-sdk-go/scw"
)

//func (s *Session) NewS3Client(t *testing.T) *S3Client {
//	return s.NewS3ClientByName(t, "default")
//}

func newS3Client(ctx context.Context, region scw.Region) *s3.Client {
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
		customEndpoint = "https://s3." + region.String() + ".scw.cloud"
	}

	return s3.New(s3.Options{
		APIOptions:    nil,
		ClientLogMode: 0,
		Credentials: aws.CredentialsProviderFunc(func(ctx context.Context) (aws.Credentials, error) {
			return aws.Credentials{
				AccessKeyID:     accessKey,
				SecretAccessKey: secretKey,
			}, nil
		}),
		EndpointResolver: s3.EndpointResolverFunc(func(region string, options s3.EndpointResolverOptions) (aws.Endpoint, error) {
			return aws.Endpoint{
				URL:               customEndpoint,
				HostnameImmutable: true,
				SigningRegion:     region,
			}, nil
		}),
		//Logger:                         logging.LoggerFunc(func(classification logging.Classification, format string, v ...interface{}) {
		//	t.Logf(format, v)
		//}),
		Region: region.String(),
		//HTTPClient:                     nil,
	})

	//cfg, err := config.LoadDefaultConfig(ctx,
	//	config.WithRegion(region.String()),
	//	config.WithCredentialsProvider(aws.CredentialsProvider()),
	//	config.WithEndpointResolverWithOptions(aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
	//		if service == s3.ServiceID {
	//
	//			return aws.Endpoint{
	//				URL:               s.Config.Endpoint(),
	//				SigningRegion:     s.Config.Region,
	//				HostnameImmutable: true,
	//			}, nil
	//
	//		}
	//
	//		return aws.Endpoint{}, fmt.Errorf("unknown endpoint requested")
	//
	//	})),
	//)

	//cfg := &aws.Config{}
	//config.WithRegion(region)
	//config.WithCredentials(credentials.NewStaticCredentials(accessKey, secretKey, ""))
	//if ep := os.Getenv("SCW_S3_ENDPOINT"); ep != "" {
	//	config.WithEndpoint(ep)
	//} else {
	//	config.WithEndpoint("https://s3." + region + ".scw.cloud")
	//}
	//config.WithHTTPClient(httpClient)
	//if logging.IsDebugOrHigher() {
	//	config.WithLogLevel(aws.LogDebugWithHTTPBody)
	//}
	//
	//s, err := session.NewSession(config)
	//if err != nil {
	//	return nil, err
	//}
	//return s3.New(s)
}
