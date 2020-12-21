// +build integration

package testing

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mleone10/endpoint/internal/api"
	mockAuth "github.com/mleone10/endpoint/internal/auth/mock"
	mockDb "github.com/mleone10/endpoint/internal/dynamo/mock"
)

var s *httptest.Server

func TestMain(m *testing.M) {
	setupServer()
	code := m.Run()
	os.Exit(code)
}

func setupServer() {
	os.Setenv("ENDPOINT_LOCAL", "true")
	db := mockDb.NewClient()
	authr, _ := mockAuth.NewAuthenticator()
	s = httptest.NewServer(api.NewServer(db, authr))
}
