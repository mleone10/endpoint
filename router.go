package main

import (
	"net/http"
)

func (s *server) routes() {
	s.router.HandleFunc("/health", s.handleHealth())
	s.router.Handle("/", http.NotFoundHandler())
}
