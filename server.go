package main

import (
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
