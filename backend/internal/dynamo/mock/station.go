package mock

import (
	"github.com/mleone10/endpoint/internal/account"
	"github.com/mleone10/endpoint/internal/station"
)

// SaveStation implements a mocked call to persist a station in Dynamo.
func (m *Client) SaveStation(id account.ID, s station.Station) error {
	return nil
}
