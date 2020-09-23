// +build integration

package testing

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mleone10/endpoint/internal/dynamo/mock"
	"github.com/mleone10/endpoint/internal/server"
)

var s *httptest.Server

func TestMain(m *testing.M) {
	setupServer()
	code := m.Run()
	os.Exit(code)
}

func setupServer() {
	s = httptest.NewServer(server.NewServer(mock.NewClient()))
}
