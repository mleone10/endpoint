package main

import (
	"log"
	"net/http"
	"os"

	"github.com/mleone10/endpoint/internal/api"
	"github.com/mleone10/endpoint/internal/dynamo"
	"github.com/mleone10/endpoint/internal/firebase"
	"github.com/mleone10/endpoint/internal/firebase/mock"
)

func main() {
	db := dynamo.NewClient()
	authr := initAuthenticator()

	log.Println("Running Endpoint server on localhost port 8080")
	log.Fatal(http.ListenAndServe(":8080", api.NewServer(db, authr)))
}

func initAuthenticator() api.Authenticator {
	var authr api.Authenticator
	authr, err := firebase.NewAuthenticator()
	if _, ok := os.LookupEnv("ENDPOINT_LOCAL"); ok {
		authr, err = mock.NewAuthenticator()
	}

	if err != nil {
		log.Panicf("Failed to initialize authenticator: %v", err)
	}
	return authr
}
