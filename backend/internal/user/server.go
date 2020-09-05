package user

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// Server is an http.Handler for user interactions
type Server struct {
	router *chi.Mux
}

// NewServer returns an initialized Server
func NewServer() *Server {
	s := &Server{
		router: chi.NewRouter(),
	}

	s.router.Use(AuthTokenVerifier())
	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.PlainText(w, r, "Hello there!")
	})

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
