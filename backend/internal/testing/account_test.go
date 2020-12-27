// +build integration

package testing_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/mleone10/endpoint/internal/account"
)

const (
	headerAuthKey   = "Authorization"
	headerAuthValue = "Bearer testIDToken"
)

func TestListAPIKeys(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts/testUserID/api-keys", s.URL), nil)
	req.Header.Set(headerAuthKey, headerAuthValue)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("failed to make GET /account/api-keys request")
	}

	if res.StatusCode != http.StatusOK {
		t.Fatalf("GET /account/api-keys returned wrong status code; got %v, wanted %v", res.Status, http.StatusOK)
	}

	defer res.Body.Close()
	ks := struct {
		Keys []account.APIKey `json:"apiKeys"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&ks)
	if err != nil {
		t.Fatal("could not parse GET /account/api-keys response")
	}

	if len(ks.Keys) != 2 {
		t.Fatalf("incorrect number of API keys; got %v, wanted 2", len(ks.Keys))
	}
}

func TestGetAPIKeysForDifferentAccount(t *testing.T) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("%s/accounts/incorrectUserID/api-keys", s.URL), nil)
	req.Header.Set(headerAuthKey, headerAuthValue)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("failed to make GET /account/api-keys request")
	}

	if res.StatusCode != http.StatusForbidden {
		t.Fatalf("GET /account/api-keys returned wrong status code; got %v, wanted %v", res.Status, http.StatusOK)
	}
}

func TestPostAPIKeys(t *testing.T) {
	req, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("%s/accounts/testUserID/api-keys", s.URL), bytes.NewBuffer([]byte(`{"readOnly": true}`)))
	req.Header.Set(headerAuthKey, headerAuthValue)

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		t.Fatal("failed to make POST /account/api-keys request")
	}

	if res.StatusCode != http.StatusNoContent {
		t.Fatalf("POST /account/api-keys returned %v status, but expected %v", res.StatusCode, http.StatusNoContent)
	}
}
