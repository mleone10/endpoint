package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/mleone10/endpoint/internal/account"
)

const (
	urlParamAccountID = "accountID"
	urlParamAPIKey    = "apiKey"
)

func getAccountID(r *http.Request) account.ID {
	return account.ID(chi.URLParam(r, urlParamAccountID))
}

func getAPIKey(r *http.Request) string {
	return chi.URLParam(r, urlParamAPIKey)
}
