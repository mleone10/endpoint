package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mleone10/endpoint/internal/account"
)

type urlParamKey string
type headerKey string

const (
	urlParamAccountID = urlParamKey("accountID")
	urlParamAPIKey    = urlParamKey("apiKey")
	urlParamStationID = urlParamKey("stationID")

	headerAPIKey        = headerKey("x-api-key")
	headerAuthorization = headerKey("Authorization")
)

func getURLParam(r *http.Request, p urlParamKey) string {
	return chi.URLParam(r, string(p))
}

func getHeader(r *http.Request, h headerKey) string {
	return r.Header.Get(string(h))
}

func getAccountID(r *http.Request) account.ID {
	return account.ID(getURLParam(r, urlParamAccountID))
}
