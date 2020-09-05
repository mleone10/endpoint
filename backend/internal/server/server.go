package server

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/mleone10/endpoint/internal/server/middleware"
	"github.com/mleone10/endpoint/internal/user"
)

// Server is a root-level http.Handler
type Server struct {
	router *chi.Mux
	logger *log.Logger
}

// NewServer returns an initialized Server
func NewServer() *Server {
	s := &Server{
		router: chi.NewRouter(),
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}

	s.router.Use(middleware.ErrorReporter())
	s.router.Get("/health", s.handleHealth())
	s.router.Mount("/user", user.NewServer())

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
