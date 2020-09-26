package mock

import "github.com/mleone10/endpoint/internal/user"

// SaveUser implements a mock call to persist a user to the database.
func (m *Client) SaveUser(u *user.User) error {
	return nil
}

// GetUser implements a mock call to retrieve a user from the database
func (m *Client) GetUser(u user.ID) (*user.User, error) {
	return &user.User{
		ID: user.ID("testUserID"),
		APIKeys: []*user.APIKey{
			&user.APIKey{Key: "key1", ReadOnly: true},
			&user.APIKey{Key: "key2", ReadOnly: false},
		},
	}, nil
}
