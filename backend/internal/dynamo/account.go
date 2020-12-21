package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mleone10/endpoint/internal/account"
)

const (
	skPrefixAPIKey = "KEY"
)

// User represents a user's basic info in the database
type User struct {
	itemKey
}

// APIKey represents an API key database item
type APIKey struct {
	itemKey
	Key      string `json:"key"`
	ReadOnly bool   `json:"readOnly"`
}

// SaveAPIKey upserts an APIKey in the Dynamo database
func (c *Client) SaveAPIKey(apiKey *account.APIKey) error {
	return nil
}

// ListAPIKeys returns a slice of APIKeys for a given account ID
func (c *Client) ListAPIKeys(id account.ID) ([]account.APIKey, error) {
	uidKey, skKey := ":uid", ":skKey"
	res, err := c.db.Query(&dynamodb.QueryInput{
		TableName:              aws.String(endpointTableName),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = %s and begins_with(%s, %s)", endpointPK, uidKey, endpointSK, skKey)),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			uidKey: {S: aws.String(id.String())},
			skKey:  {S: aws.String(skPrefixAPIKey)},
		},
	})
	if err != nil {
		return nil, ErrorInternalServerError
	}

	ks := []account.APIKey{}
	for _, i := range res.Items {
		k := account.APIKey{}
		dynamodbattribute.UnmarshalMap(i, &k)
		ks = append(ks, k)
	}

	return ks, nil
}

// // SaveUser inserts a User into the database
// func (c *Client) SaveUser(u *account.User) error {
// 	transactItems := []*dynamodb.TransactWriteItem{}
// 	for _, k := range u.APIKeys {
// 		item, err := dynamodbattribute.MarshalMap(&APIKey{
// 			itemKey: itemKey{
// 				PK: u.ID.String(),
// 				SK: fmt.Sprintf("%s#%s", skPrefixAPIKey, k.Key),
// 			},
// 			Key:      k.Key,
// 			ReadOnly: k.ReadOnly,
// 		})
// 		if err != nil {
// 			return fmt.Errorf("could not marshal APIKey to DynamoDB item")
// 		}

// 		transactItems = append(transactItems, &dynamodb.TransactWriteItem{
// 			Put: &dynamodb.Put{
// 				TableName: aws.String(endpointTableName),
// 				Item:      item,
// 			},
// 		})
// 	}

// 	user, err := dynamodbattribute.MarshalMap(&User{
// 		itemKey: itemKey{
// 			PK: u.ID.String(),
// 			SK: skPrefixUser,
// 		},
// 	})
// 	if err != nil {
// 		return fmt.Errorf("could not marshal user info to User item")
// 	}
// 	transactItems = append(transactItems, &dynamodb.TransactWriteItem{
// 		Put: &dynamodb.Put{
// 			TableName:           aws.String(endpointTableName),
// 			Item:                user,
// 			ConditionExpression: aws.String(fmt.Sprintf("attribute_not_exists(%s)", endpointSK)),
// 		},
// 	})

// 	_, err = c.db.TransactWriteItems(&dynamodb.TransactWriteItemsInput{
// 		TransactItems: transactItems,
// 	})
// 	if err != nil {
// 		if awserr, ok := err.(awserr.Error); ok && awserr.Code() == dynamodb.ErrCodeTransactionCanceledException {
// 			return ErrorConflict
// 		}
// 		return fmt.Errorf("could not persist user to database: %w", err)
// 	}

// 	return nil
// }
