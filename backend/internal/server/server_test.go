package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/mleone10/endpoint/internal/dynamo/mock"
)

func TestHealth(t *testing.T) {
	req, err := http.NewRequest(http.MethodGet, "/health", nil)
	if err != nil {
		t.Fatal("failed to create request")
	}

	rr := httptest.NewRecorder()

	server := NewServer(mock.NewMockClient())

	server.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("health endpoint returned wrong status code; got %v, wanted %v", status, http.StatusOK)
	}

	t.Log(rr.Body)
	expected := `{"api":true,"db":true}`
	if strings.TrimSpace(rr.Body.String()) != expected {
		t.Errorf("health endpoint returned wrong body; got %v, wanted %v", rr.Body.String(), expected)
	}
}
