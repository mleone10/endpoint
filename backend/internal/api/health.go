package api

import (
	"net/http"

	"github.com/go-chi/render"
)

func (s *Server) handleHealth() http.HandlerFunc {
	type health struct {
		API bool `json:"api"`
		DB  bool `json:"db"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, health{API: true, DB: s.db.Health()})
	}
}
