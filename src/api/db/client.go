package db

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

type Config struct {
	region       string
	endpoint     string
	accessKey    string
	secretKey    string
	sessionToken string
}

type Db struct {
	Client *dynamodb.Client
}

func NewConfig() Config {
	return Config{
		region:       os.Getenv("AWS_DYNAMO_REGION"),
		endpoint:     os.Getenv("AWS_DYNAMO_ENDPOINT"),
		accessKey:    os.Getenv("AWS_DYNAMO_ACCESS_KEY"),
		secretKey:    os.Getenv("AWS_DYNAMO_SECRET_KEY"),
		sessionToken: os.Getenv("AWS_DYNAMO_SESSION_TOKEN"),
	}
}

func NewDynamodb(cnf Config) Db {
	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(cnf.region),
		config.WithEndpointResolver(aws.EndpointResolverFunc(
			func(service, region string) (aws.Endpoint, error) {
				return aws.Endpoint{URL: cnf.endpoint}, nil
			})),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID:     cnf.accessKey,
				SecretAccessKey: cnf.secretKey,
				SessionToken:    cnf.sessionToken,
			},
		}),
	)

	if err != nil {
		log.Fatalf("Could not connect to dynamodb. %v", err)
	}

	client := dynamodb.NewFromConfig(cfg)

	return Db{Client: client}
}
