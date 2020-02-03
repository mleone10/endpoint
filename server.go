package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
)

type server struct {
	router *http.ServeMux
	logger *log.Logger

	stationRouter *stationRouter
}

func newServer() *server {
	s := &server{
		router: http.NewServeMux(),
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}
	s.stationRouter = newStationRouter()
	s.routes()
	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *server) parse(w http.ResponseWriter, r *http.Request, data interface{}) {
	err := json.NewDecoder(r.Body).Decode(data)
	if err != nil {
		s.respond(w, r, http.StatusBadRequest, nil)
	}
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

func (s *server) handleNotFound() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Println("No matching path found")
		s.respond(w, r, http.StatusNotFound, nil)
	}
}

func (s *server) log(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.logger.Printf("Entering %v", r.URL.Path)
		h(w, r)
		s.logger.Printf("Exiting %v", r.URL.Path)
	}
}
