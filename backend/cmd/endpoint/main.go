package main

import (
	"log"
	"net/http"

	"github.com/mleone10/endpoint/internal"
	"github.com/mleone10/endpoint/internal/dynamo"
)

func main() {
	db := dynamo.NewClient()

	log.Println("Running Endpoint server on localhost port 8080")
	log.Fatal(http.ListenAndServe(":8080", internal.NewServer(db)))
}
