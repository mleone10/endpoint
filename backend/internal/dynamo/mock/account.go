package mock

import (
	"github.com/mleone10/endpoint/internal/account"
)

// SaveAPIKey implements a mock call to persist an APIKey to the database.
func (m *Client) SaveAPIKey(id account.ID, apiKey account.APIKey) error {
	return nil
}

// ListAPIKeys implements a mock call to retrieve a given user's APIKeys from the database.
func (m *Client) ListAPIKeys(id account.ID) ([]account.APIKey, error) {
	return []account.APIKey{
		{Key: "testAPIKeyOne", ReadOnly: true},
		{Key: "testAPIKeyTwo", ReadOnly: false},
	}, nil
}

// DeleteAPIKey implements a mock call to delete a given user's APIKey.
func (m *Client) DeleteAPIKey(id account.ID, apiKey account.APIKey) error {
	return nil
}
