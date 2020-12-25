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

	gsiKeyUsers = "keyUsers"
)

// User represents a user's basic info in the database.
type User struct {
	itemKey
}

// APIKeyItem represents an API key database item.
type APIKeyItem struct {
	itemKey
	Key      string `json:"key"`
	ReadOnly bool   `json:"readOnly"`
}

// KeyUserItem represents a Key-User database item.
type KeyUserItem struct {
	itemKey
	ReadOnly bool `json:"readOnly"`
}

// SaveAPIKey upserts an APIKey in the Dynamo database.
func (c *Client) SaveAPIKey(id account.ID, apiKey account.APIKey) error {
	item, err := dynamodbattribute.MarshalMap(&APIKeyItem{
		itemKey: itemKey{
			PK: id.String(),
			SK: fmt.Sprintf("%s#%s", skPrefixAPIKey, apiKey.Key),
		},
		Key:      apiKey.Key,
		ReadOnly: apiKey.ReadOnly,
	})
	if err != nil {
		return fmt.Errorf("could not marshal APIKey to DynamoDB item: %v", err)
	}

	_, err = c.db.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String(endpointTableName),
		Item:      item,
	})
	if err != nil {
		return fmt.Errorf("failed to persist APIKey to database: %v", err)
	}

	return nil
}

// ListAPIKeys returns a slice of APIKeys for a given account ID.
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

// DeleteAPIKey removes a given APIKey from a user's database record, if it exists.
func (c *Client) DeleteAPIKey(id account.ID, apiKey account.APIKey) error {
	_, err := c.db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(endpointTableName),
		Key: map[string]*dynamodb.AttributeValue{
			endpointPK: {S: aws.String(id.String())},
			endpointSK: {S: aws.String(fmt.Sprintf("%s#%s", skPrefixAPIKey, apiKey.Key))},
		},
	})

	return err
}

// GetKeyPermission queries the database for the Permission granted by a given APIKey.
func (c *Client) GetKeyPermission(a string) (account.Permission, error) {
	if a == "" {
		return account.Permission{}, fmt.Errorf("must provide non-empty API key")
	}

	apiKeyKey := ":key"
	res, err := c.db.Query(&dynamodb.QueryInput{
		TableName:              aws.String(endpointTableName),
		IndexName:              aws.String(gsiKeyUsers),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = %s", endpointSK, apiKeyKey)),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			apiKeyKey: {S: aws.String(fmt.Sprintf("%s#%s", skPrefixAPIKey, a))},
		},
	})
	if err != nil {
		return account.Permission{}, err
	}

	if len(res.Items) == 0 {
		return account.Permission{}, ErrorItemNotFound
	}

	if len(res.Items) > 1 {
		return account.Permission{}, fmt.Errorf("more than one account ID found for given API key [%s]", a)
	}

	var item KeyUserItem
	dynamodbattribute.UnmarshalMap(res.Items[0], &item)

	return account.Permission{
		ID:       account.ID(item.SK),
		ReadOnly: item.ReadOnly,
	}, nil
}
