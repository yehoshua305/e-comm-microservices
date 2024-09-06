package db

import (
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)


type Table struct {
	TableName string
	DynamodbClient *dynamodb.Client
}
