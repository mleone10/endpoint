package server

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/mleone10/endpoint/internal/server/middleware"
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

	s.router.Use(middleware.ErrorReporter())
	s.router.Use(cors.AllowAll().Handler)
	s.router.Get("/health", s.handleHealth())
	s.router.Mount("/user", user.NewServer(s.db))

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
