package main

import (
	"net/http"
)

func (s *server) handleHealth() http.HandlerFunc {
	type response struct {
		Api bool `json:"api"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, response{Api: true})
	}
}
