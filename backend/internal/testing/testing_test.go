// +build integration

package testing

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mleone10/endpoint/internal/api"
	"github.com/mleone10/endpoint/internal/dynamo/mock"
)

var s *httptest.Server

func TestMain(m *testing.M) {
	setupServer()
	code := m.Run()
	os.Exit(code)
}

func setupServer() {
	os.Setenv("ENDPOINT_LOCAL", "true")
	s = httptest.NewServer(api.NewServer(mock.NewClient()))
}
