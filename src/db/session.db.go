package db

import (
	"context"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Create Customer Session Token
func (db  *Table) CreateSession(ctx context.Context, arg Session) (Session, error) {
	arg.PK = fmt.Sprintf("CUSTOMER#%s",arg.Username)
	arg.SK = fmt.Sprintf("SESSION#%s", arg.ID.String())

	itemMap, err := attributevalue.MarshalMap(arg)
	if err != nil {
		return Session{}, err
	}

	_, err = db.DynamodbClient.PutItem(ctx, &dynamodb.PutItemInput{
		TableName: aws.String(db.TableName),
		Item: itemMap,
	})

	return arg, nil
}

// Get Customer Session Token
func (db *Table) GetSession(ctx context.Context, username string, sessionID string) (Session, error) {
	session, err := db.DynamodbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(db.TableName),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: fmt.Sprintf("CUSTOMER#%s", username)},
			"SK": &types.AttributeValueMemberS{Value: fmt.Sprintf("SESSION#%s", sessionID)},
		},
	})

	if err != nil {
		return Session{}, err
	}

	if len(session.Item) == 0 {
		return Session{}, fmt.Errorf("session for customer %s not found", username)
	}

	returnSession := Session{}
	err = attributevalue.UnmarshalMap(session.Item, &returnSession)
	return returnSession, nil
}