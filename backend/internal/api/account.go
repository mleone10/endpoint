package api

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/mleone10/endpoint/internal/account"
)

func (s *Server) handleListAPIKeys() http.HandlerFunc {
	type res struct {
		Keys []account.APIKey `json:"apiKeys"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := idFromContext(r.Context())
		if err != nil {
			s.internalServerError(w, err)
			return
		}

		u, err := s.db.ListAPIKeys(uid)
		if err != nil {
			s.internalServerError(w, err)
			return
		}

		render.JSON(w, r, res{u})
	}
}

func (s *Server) handlePostAPIKeys() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
