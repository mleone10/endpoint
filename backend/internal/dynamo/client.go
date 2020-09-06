package dynamo

import "github.com/mleone10/endpoint/internal/user"

// Client represents a DynamoDB accessor.
type Client struct {
}

// NewClient returns an initialized Dynamo Client.
func NewClient() (*Client, error) {
	return &Client{}, nil
}

// GetAPIKeys fetches a list of API keys for a given user, or an error if they cannot be retrieved.
func (c *Client) GetAPIKeys(uid *user.ID) ([]user.APIKey, error) {
	return []user.APIKey{user.APIKey(uid.Value())}, nil
}
