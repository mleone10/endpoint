package mock

import (
	"github.com/mleone10/endpoint/internal/account"
)

// SaveAPIKey implements a mock call to persist an APIKey to the database.
func (m *Client) SaveAPIKey(apiKey *account.APIKey) error {
	return nil
}

// ListAPIKeys implements a mock call to retrieve a given user's APIKeys from the database
func (m *Client) ListAPIKeys(u account.ID) ([]account.APIKey, error) {
	return []account.APIKey{
		{ID: "testUserID", Key: "testAPIKeyOne", ReadOnly: true},
		{ID: "testUserID", Key: "testAPIKeyTwo", ReadOnly: false},
	}, nil
}
