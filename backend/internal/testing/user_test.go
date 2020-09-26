// +build integration

package testing

import (
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/mleone10/endpoint/internal/user"
)

func TestGetUser(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("%s/user", s.URL))
	if err != nil {
		t.Fatal("failed to make health check request")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("GET /user returned wrong status code; got %v, wanted %v", res.Status, http.StatusOK)
	}

	defer res.Body.Close()
	u := user.User{}
	err = json.NewDecoder(res.Body).Decode(&u)
	if err != nil {
		t.Errorf("could not parse GET /user response")
	}

	if u.ID == "" {
		t.Errorf("incorrect user ID; got nil, wanted not nil")
	}

	if len(u.APIKeys) != 2 {
		t.Errorf("incorrect number of user API keys; got %v, wanted 2", len(u.APIKeys))
	}
}
