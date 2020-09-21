package main

import (
	"log"
	"net/http"

	"github.com/mleone10/endpoint/internal/dynamo"
	"github.com/mleone10/endpoint/internal/server"
)

func main() {
	db := dynamo.NewClient()

	log.Println("Running Endpoint server on localhost port 8080")
	log.Fatal(http.ListenAndServe(":8080", server.NewServer(db)))
}
