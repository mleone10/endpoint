package testing_test

import (
	"encoding/json"
	"net/http"
	"testing"

	"github.com/mleone10/endpoint/internal/station"
	test "github.com/mleone10/endpoint/internal/testing"
)

func TestListStations(t *testing.T) {
	req := authenticatedReq(http.MethodGet, "/stations", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("request failed: %v", err)
	}

	test.AssertEquals(t, http.StatusOK, res.StatusCode)

	defer res.Body.Close()
	ss := struct {
		IDs []station.ID `json:"stations"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&ss)
	if err != nil {
		t.Error("failed to parse response")
	}

	test.AssertEquals(t, 2, len(ss.IDs))
}

func TestPostStation(t *testing.T) {
	req := authenticatedReq(http.MethodPost, "/stations", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("request failed: %v", err)
	}

	test.AssertEquals(t, http.StatusOK, res.StatusCode)

	defer res.Body.Close()
	s := struct {
		ID station.ID `json:"id"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		t.Error("failed to parse response")
	}

	test.AssertNotEquals(t, "", s.ID)
}

func TestGetStation(t *testing.T) {
	req := authenticatedReq(http.MethodGet, "/stations/stationID", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("request failed: %v", err)
	}

	test.AssertEquals(t, http.StatusOK, res.StatusCode)

	defer res.Body.Close()
	s := struct {
		ID         station.ID         `json:"id"`
		Production station.Production `json:"production"`
		Modules    []struct {
			ID   station.ID         `json:"id"`
			Type station.ModuleType `json:"type"`
		} `json:"modules"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&s)
	if err != nil {
		t.Error("failed to parse response")
	}

	test.AssertEquals(t, station.ID("stationID"), s.ID)
	test.AssertEquals(t, 1, len(s.Modules))
	test.AssertEquals(t, station.ModuleType("command"), s.Modules[0].Type)
	test.AssertNotEquals(t, "", s.Modules[0].ID)
	test.AssertEquals(t, station.Rate(5.0), s.Production[station.ResourceResearch])
}

func TestDeleteStation(t *testing.T) {
	req := authenticatedReq(http.MethodDelete, "/stations/stationID", nil)
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Errorf("request failed: %v", err)
	}

	test.AssertEquals(t, http.StatusNoContent, res.StatusCode)
}
