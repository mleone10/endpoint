// +build integration

package testing

import (
	"net/http/httptest"
	"os"
	"testing"

	"github.com/mleone10/endpoint/internal/api"
	mockDb "github.com/mleone10/endpoint/internal/dynamo/mock"
	mockFirebase "github.com/mleone10/endpoint/internal/firebase/mock"
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
	authr, _ := mockFirebase.NewAuthenticator()
	s = httptest.NewServer(api.NewServer(db, authr))
}
