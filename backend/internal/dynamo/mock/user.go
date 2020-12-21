package mock

import (
	"github.com/mleone10/endpoint/internal/account"
)

// SaveUser implements a mock call to persist a user to the database.
func (m *Client) SaveUser(u *account.User) error {
	return nil
}

// GetUser implements a mock call to retrieve a user from the database
func (m *Client) GetUser(u account.ID) (*account.User, error) {
	return &account.User{
		ID: account.ID("testUserID"),
		APIKeys: []*account.APIKey{
			{Key: "key1", ReadOnly: true},
			{Key: "key2", ReadOnly: false},
		},
	}, nil
}
