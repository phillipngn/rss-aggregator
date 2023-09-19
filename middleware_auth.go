package main

import (
	"net/http"

	"github.com/phillipngn/rss-aggregator/internal/auth"
	"github.com/phillipngn/rss-aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetApiKey(r.Header)
		if err != nil {
			responseWithError(w, http.StatusUnauthorized, "Could not authorize your identity")
			return
		}

		user, err := apiCfg.DB.GetUserByApiKey(r.Context(), apiKey)
		if err != nil {
			responseWithError(w, http.StatusNotFound, "Could not get the user from given api key")
			return
		}
		handler(w, r, user)
	}
}
