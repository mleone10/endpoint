package mock

// Client is a mock datastore client.
type Client struct {
}

// NewClient constructs and returns an initialized Client.
func NewClient() *Client {
	return &Client{}
}

// Health returns the health of the mock Client
func (m *Client) Health() bool {
	return true
}
