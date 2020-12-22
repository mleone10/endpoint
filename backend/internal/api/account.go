package api

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/mleone10/endpoint/internal/account"
)

type accountDatastore interface {
	SaveAPIKey(account.ID, *account.APIKey) error
	ListAPIKeys(account.ID) ([]account.APIKey, error)
	DeleteAPIKey(account.ID, *account.APIKey) error
}

func (s *Server) handleListAPIKeys() http.HandlerFunc {
	type res struct {
		Keys []account.APIKey `json:"apiKeys"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		uid := getAccountID(r)

		u, err := s.db.ListAPIKeys(uid)
		if err != nil {
			s.internalServerError(w, err)
			return
		}

		render.JSON(w, r, res{u})
	}
}

func (s *Server) handlePostAPIKeys() http.HandlerFunc {
	type req struct {
		ReadOnly bool `json:"readOnly"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		uid := getAccountID(r)

		var req req
		render.DecodeJSON(r.Body, &req)

		k := account.NewAPIKey(req.ReadOnly)
		err := s.db.SaveAPIKey(uid, k)
		if err != nil {
			s.internalServerError(w, err)
			return
		}

		render.NoContent(w, r)
	}
}

func (s *Server) handleDeleteAPIKey() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}

func (s *Server) handleDeleteAPIKeys() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

	}
}
