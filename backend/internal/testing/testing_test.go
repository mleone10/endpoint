// +build integration

package testing_test

import (
	"fmt"
	"io"
	"net/http"
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

func authenticatedReq(method, path string, body io.Reader) *http.Request {
	req, _ := http.NewRequest(method, fmt.Sprintf("%s%s", s.URL, path), body)
	req.Header.Set("x-api-key", "whatever")

	return req
}
