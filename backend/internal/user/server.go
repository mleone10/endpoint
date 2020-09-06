package user

import (
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Server is an http.Handler for user interactions
type Server struct {
	router *chi.Mux
	db     Datastore
}

// Datastore describes a persistence layer for user details
type Datastore interface {
	apiKeyDatastore
}

// NewServer returns an initialized Server
func NewServer(db Datastore) *Server {
	s := &Server{
		router: chi.NewRouter(),
		db:     db,
	}

	_, isLocal := os.LookupEnv("ENDPOINT_LOCAL")

	s.router.Use(OrMiddleware(isLocal, AuthStubber(), AuthTokenVerifier()))
	s.router.Get("/api-keys", func(w http.ResponseWriter, r *http.Request) {
		apiKeys, err := getAPIKeys(r.Context(), s.db)
		if err != nil {
			http.Error(w, "could not retrieve API keys", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, struct {
			Keys []APIKey `json:"keys"`
		}{
			Keys: apiKeys,
		})
	})
	s.router.Post("/api-keys", func(w http.ResponseWriter, r *http.Request) {
		req := struct {
			Nickname string `json:"nickname"`
			ReadOnly bool   `json:"readOnly"`
		}{}

		err := render.DecodeJSON(r.Body, &req)
		if err != nil {
			http.Error(w, "could not parse request", http.StatusBadRequest)
			return
		}
		if req.Nickname == "" {
			http.Error(w, "nickname must not be empty", http.StatusBadRequest)
			return
		}

		apiKey, err := newAPIKey(r.Context(), s.db, req.Nickname, req.ReadOnly)
		if err != nil {
			http.Error(w, "failed to create new API key", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, apiKey)
	})

	// TODO: Implement DELETE /api-keys/{key} (delete the given APIKey)

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
