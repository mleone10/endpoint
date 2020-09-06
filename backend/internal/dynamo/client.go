package dynamo

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/mleone10/endpoint/internal/user"
)

/*
	Table schema (PK: userID:string, SK: sortKey:string)
	API-Key schema (SK: KEY#key:string, Attrs: apiKey:string, name:string, readOnly:string)
	GSI (PK: userID:string, SK: apiKey:string)
*/

const endpointTableName = "endpoint"

// Client represents a DynamoDB accessor.
type Client struct {
	db *dynamodb.DynamoDB
}

// NewClient returns an initialized Dynamo Client.
func NewClient() (*Client, error) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		SharedConfigState: session.SharedConfigEnable,
	}))

	return &Client{
		db: dynamodb.New(sess),
	}, nil
}

// GetAPIKeys fetches a list of API keys for a given user, or an error if they cannot be retrieved.
func (c *Client) GetAPIKeys(uid *user.ID) ([]user.APIKey, error) {
	// TODO: Call DynamoDB table to retrieve keys
	return []user.APIKey{user.APIKey(uid.Value())}, nil
}
