package api

import (
	"net/http"

	"github.com/go-chi/render"
	"github.com/mleone10/endpoint/internal/account"
	"github.com/mleone10/endpoint/internal/station"
)

type stationDatastore interface {
	SaveStation(account.ID, station.Station) error
}

func (s *Server) handlePostStation() http.HandlerFunc {
	type res struct {
		ID station.ID `json:"id"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		station := station.New()
		perm, _ := getCtxPermission(r)
		err := s.db.SaveStation(perm.ID, station)
		if err != nil {
			s.internalServerError(w, err)
		}

		render.JSON(w, r, res{ID: station.ID})
	}
}
