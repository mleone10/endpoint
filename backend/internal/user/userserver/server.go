package userserver

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	"github.com/go-chi/render"
	"github.com/mleone10/endpoint/internal/dynamo"
	"github.com/mleone10/endpoint/internal/user"
)

// Server is an http.Handler for user interactions
type Server struct {
	router *chi.Mux
	logger *log.Logger
	db     user.Datastore
}

// NewServer returns an initialized Server
func NewServer(logger *log.Logger, db user.Datastore) *Server {
	s := &Server{
		router: chi.NewRouter(),
		logger: logger,
		db:     db,
	}

	_, isLocal := os.LookupEnv("ENDPOINT_LOCAL")

	s.router.Use(OrMiddleware(isLocal, AuthStubber(), AuthTokenVerifier()))
	s.router.Get("/", s.handleGetUser())
	s.router.Post("/", s.handlePostUser())

	return s
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.router.ServeHTTP(w, r)
}

func (s *Server) handleGetUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := IDFromContext(r.Context())
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		u, err := s.db.GetUser(uid)
		if err == dynamo.ErrorItemNotFound {
			http.NotFound(w, r)
			return
		} else if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}

		render.JSON(w, r, u)
	}
}

func (s *Server) handlePostUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := IDFromContext(r.Context())
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
			return
		}
		u := user.New(uid)

		err = s.db.SaveUser(u)
		if err != nil {
			log.Printf("could not save user to database: %v", err)
			if err == dynamo.ErrorConflict {
				http.Error(w, "user already exists", http.StatusConflict)
				return
			} else {
				http.Error(w, "internal server error", http.StatusInternalServerError)
				return
			}
		}

		render.JSON(w, r, u)
	}
}
