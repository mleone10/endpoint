package main

import (
	"log"
	"net/http"

	"github.com/mleone10/endpoint/internal"
)

func main() {
	log.Println("Running Endpoint server on localhost port 8080")
	log.Fatal(http.ListenAndServe(":8080", internal.NewServer()))
}
