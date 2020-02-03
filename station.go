package main

import (
	"net/http"
)

type stationRouter struct {
	*http.ServeMux
	*server
}

type station struct {
	id   id
	name string
}

type stationRequest struct {
	Name string `json:"name"`
}
type stationResponse struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func newStationRouter() *stationRouter {
	s := &stationRouter{
		ServeMux: http.NewServeMux(),
	}
	s.routes()
	return s
}

func (s *stationRouter) routes() {
	s.HandleFunc("/", s.handleStation())
}

func (s *stationRouter) handle() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.ServeMux.ServeHTTP(w, r)
	}
}

func (s *server) handleStations() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		s.respond(w, r, http.StatusOK, "Many stations")
	}
}

func (s *stationRouter) handleStation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case MethodPost:
			req := stationRequest{}
			s.parse(w, r, &req)

			station := newStation().withName(req.Name).save()

			res := stationResponse{
				Id:   station.id.value,
				Name: station.name,
			}
			s.respond(w, r, http.StatusOK, &res)
		default:
			s.respond(w, r, http.StatusOK, "One station")
		}
	}
}

func newStation() *station {
	return &station{
		id: newId(),
	}
}

func (s *station) withName(name string) *station {
	s.name = name
	return s
}

func (s *station) save() *station {
	return s
}
