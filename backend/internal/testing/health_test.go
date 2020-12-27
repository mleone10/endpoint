// +build integration

package testing_test

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHealth(t *testing.T) {
	res, err := http.Get(fmt.Sprintf("%s/health", s.URL))
	if err != nil {
		t.Fatal("failed to make health check request")
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("health check returned wrong status code; got %v, wanted %v", res.Status, http.StatusOK)
	}

}
