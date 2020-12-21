// +build integration

package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/mleone10/endpoint/internal/account"
)

func TestListAPIKeys(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("%s/account/api-keys", s.URL))
	if err != nil {
		t.Fatal("failed to make GET /account/api-keys request")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("GET /account/api-keys returned wrong status code; got %v, wanted %v", res.Status, http.StatusOK)
	}

	defer res.Body.Close()
	ks := struct {
		Keys []account.APIKey `json:"apiKeys"`
	}{}
	err = json.NewDecoder(res.Body).Decode(&ks)
	if err != nil {
		t.Errorf("could not parse GET /account/api-keys response")
	}

	if len(ks.Keys) != 2 {
		t.Errorf("incorrect number of API keys; got %v, wanted 2", len(ks.Keys))
	}
}
