package main

import (
	"log"
	"net/http"

	"github.com/mleone10/endpoint/internal/dynamo"
	"github.com/mleone10/endpoint/internal/server"
)

func main() {
	db, err := dynamo.NewClient()
	if err != nil {
		log.Fatalf("Error initializing database: %w", err)
	}

	log.Println("Running Endpoint server on localhost port 8080")
	log.Fatal(http.ListenAndServe(":8080", server.NewServer(db)))
}
