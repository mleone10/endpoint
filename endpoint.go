package main

import (
	"log"
	"net/http"
)

func main() {
	log.Println("Running Endpoint server on localhost port 8080")
	log.Fatal(http.ListenAndServe(":8080", newServer()))
}
