package dynamo

import "github.com/mleone10/endpoint/internal/user"

// Client represents a DynamoDB accessor.
type Client struct {
}

// NewClient returns an initialized Dynamo Client.
func NewClient() (*Client, error) {
	// TODO: Initialize underlying DynamoDB client
	return &Client{}, nil
}

// GetAPIKeys fetches a list of API keys for a given user, or an error if they cannot be retrieved.
func (c *Client) GetAPIKeys(uid *user.ID) ([]user.APIKey, error) {
	// TODO: Create DynamoDB table using Serverless framework
	/*
		Table schema (PK: userID:string, SK: sortKey:string)
		API-Key schema (SK: KEY#key:string, Attrs: apiKey:string, name:string, readOnly:string)
		GSI (PK: userID:string, SK: apiKey:string)
	*/
	// TODO: Call DynamoDB table to retrieve keys
	return []user.APIKey{user.APIKey(uid.Value())}, nil
}
