package db

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/yehoshua305/e-comm-microservices/src/util"
)

type Table struct {
	TableName string
	DynamodbClient *dynamodb.Client
}

var testDB *Table

func TestMain(m *testing.M) {
	configVariables, err := util.LoadConfig("../.")
	if err != nil {
		log.Fatalf("cannot load config %v", err)
	}
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("eu-west-1"),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: configVariables.DYNAMODB_ACCESS_KEYID, SecretAccessKey: configVariables.DYNAMODB_SECRET_ACCESS_KEY, SessionToken: "",
				Source: "Hard coded credentials for local DynamoDB",
			},
		}),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

    // Create a DynamoDB client
    dynamoDBClient := dynamodb.NewFromConfig(cfg, func (o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://localhost:8000/")	
	})

	testDB = &Table{TableName: "Test", DynamodbClient: dynamoDBClient}
	os.Exit(m.Run())
}