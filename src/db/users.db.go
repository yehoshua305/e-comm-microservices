package db

import (
	"context"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

// Create User item
func createItem(arg User) (map[string]types.AttributeValue, map[string]types.AttributeValue, error) {
	arg.PK = fmt.Sprintf("User#%s", arg.Username)
	arg.SK = fmt.Sprintf("User#%s", arg.Username)
	arg.CreatedAt = time.Now()

	item1Map, err1 := attributevalue.MarshalMap(arg)
	if err1 != nil {
		return nil, nil, err1
	}

	item2 := struct {
		PK       string `dynamodbav:"PK"`
		SK       string `dynamodbav:"SK"`
		Username string `dynamodbav:"Username"`
		Email    string `dynamodbav:"Email"`
	}{
		PK:       fmt.Sprintf("User#%s", arg.Email),
		SK:       fmt.Sprintf("User#%s", arg.Email),
		Username: arg.Username,
		Email:    arg.Email,
	}
	item2Map, err2 := attributevalue.MarshalMap(item2)
	if err2 != nil {
		return nil, nil, err2
	}

	return item1Map, item2Map, nil

}

// Create User
func (db *Table) CreateUser(ctx context.Context, arg User) (User, error) {

	item1, item2, err := createItem(arg)
	if err != nil {
		return User{}, fmt.Errorf("error creating User %w", err)
	}

	_, err = db.DynamodbClient.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: []types.TransactWriteItem{
			{
				Put: &types.Put{
					TableName:           aws.String(db.TableName),
					Item:                item1,
					ConditionExpression: aws.String("attribute_not_exists(PK)"),
				},
			},
			{
				Put: &types.Put{
					TableName:           aws.String(db.TableName),
					Item:                item2,
					ConditionExpression: aws.String("attribute_not_exists(PK)"),
				},
			},
		},
	})

	if err != nil {
		return User{}, fmt.Errorf("error writing User to db %w", err)
	}

	var User User
	err = attributevalue.UnmarshalMap(item1, &User)
	if err != nil {
		return User, fmt.Errorf("error unmarshalling User, %w", err)
	}
	return User, nil
}

// Get User
func (db *Table) GetUser(ctx context.Context, username string) (User, error) {
	result, err := db.DynamodbClient.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(db.TableName),
		Key: map[string]types.AttributeValue{
			"PK": &types.AttributeValueMemberS{Value: fmt.Sprintf("User#%s", username)},
			"SK": &types.AttributeValueMemberS{Value: fmt.Sprintf("User#%s", username)},
		},
	})
	if err != nil {
		return User{}, err
	}

	if len(result.Item) == 0 {
		return User{}, fmt.Errorf("User with username %s not found", username)
	}

	User := User{}
	err = attributevalue.UnmarshalMap(result.Item, &User)
	return User, err
}

// Update User

// Delete User
func (db *Table) DeleteUser(ctx context.Context, username string) (string, error) {
	User, err := db.GetUser(ctx, username)
	if err != nil {
		return fmt.Sprintf("Error getting User %s", username), err
	}
	_, err = db.DynamodbClient.TransactWriteItems(ctx, &dynamodb.TransactWriteItemsInput{
		TransactItems: []types.TransactWriteItem{
			{
				Delete: &types.Delete{
					TableName: aws.String(db.TableName),
					Key: map[string]types.AttributeValue{
						"PK": &types.AttributeValueMemberS{Value: fmt.Sprintf("User#%s", username)},
						"SK": &types.AttributeValueMemberS{Value: fmt.Sprintf("User#%s", username)},
					},
				},
			},
			{
				Delete: &types.Delete{
					TableName: aws.String(db.TableName),
					Key: map[string]types.AttributeValue{
						"PK": &types.AttributeValueMemberS{Value: fmt.Sprintf("User#%s", User.Email)},
						"SK": &types.AttributeValueMemberS{Value: fmt.Sprintf("User#%s", User.Email)},
					},
				},
			},
		},
	})

	return fmt.Sprintf("User %s deleted", username), err
}
