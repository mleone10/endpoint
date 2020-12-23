package api

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
)

// Server is a root-level http.Handler
type Server struct {
	router *chi.Mux
	logger *log.Logger
	db     Datastore
}

// Datastore describes the persistence operations required by the HTTP server
type Datastore interface {
	accountDatastore
	Health() bool
}

// NewServer returns an initialized Server
func NewServer(db Datastore, authr Authenticator) *Server {
	s := &Server{
		router: chi.NewRouter(),
		logger: log.New(os.Stderr, "", log.LstdFlags),
		db:     db,
	}

	s.router.Use(cors.AllowAll().Handler)
	s.router.Get("/health", s.handleHealth())
	s.router.Route(fmt.Sprintf("/accounts/{%s}", urlParamAccountID), func(r chi.Router) {
		r.Use(s.authTokenVerifier(authr))
		r.Route("/api-keys", func(r chi.Router) {
			r.Get("/", s.handleListAPIKeys())
			r.Post("/", s.handlePostAPIKeys())
			r.Delete("/", s.handleDeleteAPIKeys())
			r.Route(fmt.Sprintf("/{%s}", urlParamAPIKey), func(r chi.Router) {
				r.Delete("/", s.handleDeleteAPIKey())
			})
		})
	})
	s.router.Route("/stations", func(r chi.Router) {
		r.Use(s.keyTokenVerifier())
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			render.NoContent(w, r)
		})
	})

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
