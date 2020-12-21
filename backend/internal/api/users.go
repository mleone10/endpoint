package api

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/mleone10/endpoint/internal/dynamo"
	"github.com/mleone10/endpoint/internal/user"
)

func (s *Server) handleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := idFromContext(r.Context())
		if err != nil {
			s.internalServerError(w, err)
			return
		}

		u, err := s.db.GetUser(uid)
		if err == dynamo.ErrorItemNotFound {
			render.JSON(w, r, user.NewUser(uid))
			return
		} else if err != nil {
			s.internalServerError(w, err)
			return
		}

		render.JSON(w, r, u)
	}
}
