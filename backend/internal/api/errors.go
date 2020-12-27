package api

import "net/http"

func (s *Server) internalServerError(w http.ResponseWriter, err error) {
	s.logger.Println(err)
	http.Error(w, "internal server error", http.StatusInternalServerError)
	return
}

func (s *Server) forbidden(w http.ResponseWriter) {
	http.Error(w, "operation forbidden with given api key", http.StatusForbidden)
	return
}

func (s *Server) notFound(w http.ResponseWriter, err error) {
	s.logger.Println(err)
	http.Error(w, "resource not found", http.StatusNotFound)
	return
}
