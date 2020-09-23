package mock

import "github.com/mleone10/endpoint/internal/user"

type Client struct {
}

func NewMockClient() *Client {
	return &Client{}
}

func (m *Client) Health() bool {
	return true
}

func (m *Client) GetAPIKeys(uid *user.ID) ([]user.APIKey, error) {
	return []user.APIKey{user.APIKey{}}, nil
}

func (m *Client) PutAPIKey(uid *user.ID, apiKey *user.APIKey) error {
	return nil
}

func (m *Client) DeleteAPIKey(uid *user.ID, apiKEy *user.APIKey) error {
	return nil
}
