package mock

import (
	"github.com/mleone10/endpoint/internal/account"
	"github.com/mleone10/endpoint/internal/dynamo"
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

// GetStation implements a mock call to retrieve a station.
func (m *Client) GetStation(uid account.ID, sid station.ID) (station.Station, error) {
	if sid == station.ID("stationID") {
		s := station.New()
		s.ID = sid
		return s, nil
	}
	return station.New(), dynamo.ErrorItemNotFound
}

// DeleteStation implements a mock call to delete a station.
func (m *Client) DeleteStation(uid account.ID, sid station.ID) error {
	return nil
}
