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
	UserID   string `json:"userID"`
	SortKey  string `json:"sortKey"`
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
