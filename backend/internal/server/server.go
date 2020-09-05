package server

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mleone10/endpoint/internal/server/middleware"
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
	s.router.Group(func(r chi.Router) {
		r.Use(middleware.AuthTokenVerifier())
		r.Get("/private", func(w http.ResponseWriter, r *http.Request) {
			render.PlainText(w, r, "Hello there!")
		})
	})

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
