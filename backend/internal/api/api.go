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
func NewServer(db Datastore) *Server {
	s := &Server{
		router: chi.NewRouter(),
		logger: log.New(os.Stderr, "", log.LstdFlags),
		db:     db,
	}

	_, isLocal := os.LookupEnv("ENDPOINT_LOCAL")

	s.router.Use(cors.AllowAll().Handler)
	s.router.Get("/health", s.handleHealth())
	s.router.Route("/user", func(r chi.Router) {
		r.Use(OrMiddleware(isLocal, AuthStubber(), AuthTokenVerifier()))
		r.Get("/", s.handleGetUser())
		r.Post("/", s.handlePostUser())
	})
	s.router.Route("/stations", func(r chi.Router) {
		r.Use(KeyTokenVerifier())
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			render.NoContent(w, r)
		})
	})

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
