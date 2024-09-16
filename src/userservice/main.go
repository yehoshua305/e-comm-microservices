package main

import (
	"context"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	db "github.com/yehoshua305/e-comm-microservices/src/db"
	"github.com/yehoshua305/e-comm-microservices/src/util"
	"github.com/yehoshua305/e-comm-microservices/src/user"
)

func main() {
	configVariables, err := util.LoadConfig(".")
	if err != nil {
		log.Fatalf("cannot load config %v", err)
	}
	cfg, err := config.LoadDefaultConfig(
		context.TODO(),
		config.WithRegion("eu-west-1"),
		config.WithCredentialsProvider(credentials.StaticCredentialsProvider{
			Value: aws.Credentials{
				AccessKeyID: configVariables.DYNAMODB_ACCESS_KEYID, SecretAccessKey: configVariables.DYNAMODB_SECRET_ACCESS_KEY, SessionToken: "",
				Source: "Credentials from env file",
			},
		}),
	)

	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}

	// Create a DynamoDB client
	dynamoDBClient := dynamodb.NewFromConfig(cfg, func(o *dynamodb.Options) {
		o.BaseEndpoint = aws.String("http://localhost:8000/")
	})

	table := db.Table{TableName: "Test", DynamodbClient: dynamoDBClient}
	server, err := user.NewServer(configVariables, table)
	if err != nil {
		log.Fatalf("Error creating server: %v", err)
	}

	err = server.Start(configVariables.ServerAddress)
	if err != nil {
		log.Fatalf("Error starting server: %v", err)
	}
}
