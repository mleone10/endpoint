package api

import (
	"fmt"
	"net/http"

	"github.com/go-chi/render"
	"github.com/mleone10/endpoint/internal/account"
	"github.com/mleone10/endpoint/internal/dynamo"
	"github.com/mleone10/endpoint/internal/station"
)

type stationDatastore interface {
	SaveStation(account.ID, station.Station) error
	ListStations(account.ID) ([]station.ID, error)
	GetStation(account.ID, station.ID) (station.Station, error)
	DeleteStation(account.ID, station.ID) error
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
			return
		}

		render.JSON(w, r, res{ID: station.ID})
	}
}

func (s *Server) handleListStations() http.HandlerFunc {
	type res struct {
		Stations []station.ID `json:"stations"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		perm, _ := getCtxPermission(r)
		ss, err := s.db.ListStations(perm.ID)
		if err != nil {
			s.internalServerError(w, err)
			return
		}

		render.JSON(w, r, res{Stations: ss})
	}
}

func (s *Server) handleGetStation() http.HandlerFunc {
	type module struct {
		ID   station.ID         `json:"id"`
		Type station.ModuleType `json:"type"`
	}
	type res struct {
		ID      station.ID `json:"id"`
		Modules []module   `json:"modules"`
	}
	return func(w http.ResponseWriter, r *http.Request) {
		sid := getURLParam(r, urlParamStationID)
		perm, ok := getCtxPermission(r)
		if !ok {
			s.internalServerError(w, fmt.Errorf("station id not found"))
			return
		}

		station, err := s.db.GetStation(perm.ID, station.ID(sid))
		if err != nil && err == dynamo.ErrorItemNotFound {
			s.notFound(w, fmt.Errorf("station [%v] not found", sid))
		} else if err != nil {
			s.internalServerError(w, fmt.Errorf("failed to retrieve station [%v] for account [%v]: %v", perm.ID, sid, err))
			return
		}

		ms := []module{}
		for _, m := range station.Modules {
			ms = append(ms, module{ID: m.ID, Type: m.Type})
		}
		render.JSON(w, r, res{
			ID:      station.ID,
			Modules: ms,
		})
	}
}

func (s *Server) handleDeleteStation() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		sid := getURLParam(r, urlParamStationID)
		perm, ok := getCtxPermission(r)
		if !ok {
			s.internalServerError(w, fmt.Errorf("station id not found"))
		}

		err := s.db.DeleteStation(perm.ID, station.ID(sid))
		if err != nil {
			s.internalServerError(w, fmt.Errorf("failed to delete station [%v] for account [%v]: %v", perm.ID, sid, err))
		}

		render.NoContent(w, r)
	}
}
