package dynamo

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/mleone10/endpoint/internal/user"
)

// APIKey represents an API key database item
type APIKey struct {
	UserID   string `json:"pk"`
	SortKey  string `json:"sk"`
	Key      string `json:"key"`
	Nickname string `json:"nickname"`
	ReadOnly bool   `json:"readOnly"`
}

// GetAPIKeys fetches a list of API keys for a given user, or an error if they cannot be retrieved.
func (c *Client) GetAPIKeys(uid *user.ID) ([]user.APIKey, error) {
	res, err := c.db.Query(&dynamodb.QueryInput{
		TableName:              aws.String(endpointTableName),
		KeyConditionExpression: aws.String(fmt.Sprintf("%s = :userID", endpointPK)),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":userID": {
				S: aws.String(uid.Value()),
			},
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve API keys: %w", err)
	}

	apiKeys := []APIKey{}
	dynamodbattribute.UnmarshalListOfMaps(res.Items, &apiKeys)

	ret := []user.APIKey{}
	for _, item := range apiKeys {
		ret = append(ret, user.APIKey{
			Key:      item.Key,
			Nickname: item.Nickname,
			ReadOnly: item.ReadOnly,
		})
	}

	return ret, nil
}

// PutAPIKey writes the given APIKey to the database, or returns an error.
func (c *Client) PutAPIKey(uid *user.ID, apiKey *user.APIKey) error {
	item, err := dynamodbattribute.MarshalMap(APIKey{
		UserID:   uid.Value(),
		SortKey:  fmt.Sprintf("KEY#%s", apiKey.Key),
		Key:      apiKey.Key,
		Nickname: apiKey.Nickname,
		ReadOnly: apiKey.ReadOnly,
	})
	if err != nil {
		return fmt.Errorf("could not convert APIKey to database item: %w", err)
	}

	_, err = c.db.PutItem(&dynamodb.PutItemInput{
		TableName:           aws.String(endpointTableName),
		Item:                item,
		ConditionExpression: aws.String(fmt.Sprintf("attribute_not_exists(%s) and attribute_not_exists(%s)", endpointPK, endpointSK)),
	})
	if err != nil {
		return fmt.Errorf("failed to insert API key from database: %w", err)
	}

	return nil
}

// DeleteAPIKey removes the given APIKey from the database, or returns an error.
func (c *Client) DeleteAPIKey(uid *user.ID, apiKey *user.APIKey) error {
	key, err := dynamodbattribute.MarshalMap(struct {
		UserID  string `json:"pk"`
		SortKey string `json:"sk"`
	}{
		UserID:  uid.Value(),
		SortKey: fmt.Sprintf("KEY#%s", apiKey.Key),
	})
	if err != nil {
		return fmt.Errorf("could not create APIKey key: %w", err)
	}

	_, err = c.db.DeleteItem(&dynamodb.DeleteItemInput{
		TableName: aws.String(endpointTableName),
		Key:       key,
	})
	if err != nil {
		return fmt.Errorf("failed to delete API key from database: %w", err)
	}

	return nil
}
