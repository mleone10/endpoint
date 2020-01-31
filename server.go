package main

import (
	"encoding/json"
	"net/http"
)

type server struct {
	router *http.ServeMux
}

func newServer() *server {
	s := &server{
		router: http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) respond(w http.ResponseWriter, r *http.Request, status int, data interface{}) {
	w.WriteHeader(status)
	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			s.respond(w, r, http.StatusInternalServerError, nil)
		}
	}
}
