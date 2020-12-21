package api

import (
	"log"
	"net/http"

	"github.com/go-chi/render"
	"github.com/mleone10/endpoint/internal/dynamo"
	"github.com/mleone10/endpoint/internal/user"
)

func (s *Server) handleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := IDFromContext(r.Context())
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		u, err := s.db.GetUser(uid)
		if err == dynamo.ErrorItemNotFound {
			http.NotFound(w, r)
			return
		} else if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, u)
	}
}

func (s *Server) handlePostUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := IDFromContext(r.Context())
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		u := user.New(uid)

		err = s.db.SaveUser(u)
		if err != nil {
			log.Printf("could not save user to database: %v", err)
			if err == dynamo.ErrorConflict {
				http.Error(w, "user already exists", http.StatusConflict)
				return
			}
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, u)
	}
}
