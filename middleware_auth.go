package main

import (
	"fmt"
	"net/http"

	"github.com/haroonalbar/rss-aggregater/auth"
	"github.com/haroonalbar/rss-aggregater/internal/database"
)

// handler with database.User
type authHandler func(http.ResponseWriter, *http.Request, database.User)

// takes a authHandler and returns a HandlerFunc
func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		apiKey, err := auth.GetAPIKey(r.Header)
		if err != nil {
			respondWithJSON(w, 403, fmt.Sprintf("Auth error: %v", err))
			return
		}

		user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
		if err != nil {
			respondWithJSON(w, 400, fmt.Sprintf("Couldn't get user: %v", err))
			return
		}

		handler(w, r, user)
	}
}
