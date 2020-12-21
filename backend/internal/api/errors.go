package api

import "net/http"

func (s *Server) internalServerError(w http.ResponseWriter, err error) {
	s.logger.Println(err)
	http.Error(w, "internal server error", http.StatusInternalServerError)
	return
}
