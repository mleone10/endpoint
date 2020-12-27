package mock

import (
	"github.com/mleone10/endpoint/internal/account"
	"github.com/mleone10/endpoint/internal/station"
)

// SaveStation implements a mocked call to persist a station in Dynamo.
func (m *Client) SaveStation(id account.ID, s station.Station) error {
	return nil
}

// ListStations implements a mocked call to list all of an account's stations.
func (m *Client) ListStations(id account.ID) ([]station.ID, error) {
	return []station.ID{
		"stationIDOne",
		"stationIDTwo",
	}, nil
}
