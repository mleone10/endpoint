package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/mleone10/endpoint/internal/account"
)

type accountDatastore interface {
	SaveAPIKey(account.ID, account.APIKey) error
	ListAPIKeys(account.ID) ([]account.APIKey, error)
	DeleteAPIKey(account.ID, account.APIKey) error
}

func (s *Server) handleListAPIKeys() http.HandlerFunc {
	type res struct {
		Keys []account.APIKey `json:"apiKeys"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		accountID := getAccountID(r)

		ks, err := s.db.ListAPIKeys(accountID)
		if err != nil {
			s.internalServerError(w, err)
			return
		}

		render.JSON(w, r, res{ks})
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
		accountID := getAccountID(r)
		apiKey := getAPIKey(r)

		err := s.db.DeleteAPIKey(accountID, account.APIKey{Key: apiKey})
		if err != nil {
			s.internalServerError(w, err)
			return
		}

		render.NoContent(w, r)
	}
}

func (s *Server) handleDeleteAPIKeys() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		accountID := getAccountID(r)

		ks, err := s.db.ListAPIKeys(accountID)
		if err != nil {
			s.internalServerError(w, err)
			return
		}

		var failed bool
		for _, k := range ks {
			err = s.db.DeleteAPIKey(accountID, k)
			if err != nil {
				s.internalServerError(w, err)
			}
		}

		if failed {
			s.internalServerError(w, fmt.Errorf("failed to delete one or more API keys"))
			return
		}

		render.NoContent(w, r)
	}
}
