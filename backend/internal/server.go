package internal

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/mleone10/endpoint/internal/station"
	"github.com/mleone10/endpoint/internal/user"
	"github.com/mleone10/endpoint/internal/user/userserver"
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

	s.router.Use(cors.AllowAll().Handler)
	s.router.Get("/health", s.handleHealth())
	s.router.Mount("/user", userserver.NewServer(s.logger, s.db))
	s.router.Mount("/stations", station.NewServer())

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleHealth() http.HandlerFunc {
	type health struct {
		API bool `json:"api"`
		DB  bool `json:"db"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, health{API: true, DB: s.db.Health()})
	}
}
