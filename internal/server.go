package internal

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
)

type server struct {
	router *chi.Mux
	logger *log.Logger
}

func NewServer() *server {
	s := &server{
		router: chi.NewRouter(),
		logger: log.New(os.Stderr, "", log.LstdFlags),
	}

	s.router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello world"))
	})

	return s
}

func (s *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}
