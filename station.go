package main

import (
	"net/http"
)

func (s *server) handleStations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, "Many stations")
	}
}

func (s *server) handleStation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, "One station")
	}
}
