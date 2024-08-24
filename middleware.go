package main

import (
	"fmt"
	"net/http"

	"github.com/KatragaddaLokesh/RSSagg/auth"
	"github.com/KatragaddaLokesh/RSSagg/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apikey, err := auth.GetApiKey(r.Header)
		if err != nil {
			responwithError(w, 403, fmt.Sprintf("Auth Error %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apikey)
		if err != nil {
			responwithError(w, 400, fmt.Sprintf("Couldn't Get user %v", err))
			return
		}

		handler(w, r, user)

	}
}
