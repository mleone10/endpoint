package api

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/mleone10/endpoint/internal/user"
)

// Server is a root-level http.Handler
type Server struct {
	router *chi.Mux
	logger *log.Logger
	db     Datastore
}

// Datastore describes the persistence operations required by the HTTP server
type Datastore interface {
	user.Datastore
	Health() bool
}

// NewServer returns an initialized Server
func NewServer(db Datastore, authr authenticator) *Server {
	s := &Server{
		router: chi.NewRouter(),
		logger: log.New(os.Stderr, "", log.LstdFlags),
		db:     db,
	}

	s.router.Use(cors.AllowAll().Handler)
	s.router.Get("/health", s.handleHealth())
	s.router.Route("/user", func(r chi.Router) {
		r.Use(authTokenVerifier(authr))
		r.Get("/", s.handleGetUser())
		r.Get("/api-keys", s.handleGetAPIKeys())
		r.Post("/api-keys", s.handlePostAPIKeys())
	})
	s.router.Route("/stations", func(r chi.Router) {
		r.Use(keyTokenVerifier())
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			render.NoContent(w, r)
		})
	})

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
