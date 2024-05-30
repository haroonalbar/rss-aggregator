package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/haroonalbar/rss-aggregater/internal/database"
)

// make the handler a method to take in the apiConfig
func (apiCfg *apiConfig) handlerCreateUser(w http.ResponseWriter, r *http.Request) {
	// we are expecing a parameter
	// and it should contain a name
	type Parameters struct {
		Name string `json:"name"`
	}

	params := Parameters{}

	// create a new decoder for response body
	decoder := json.NewDecoder(r.Body)

	// decode it to params
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error decoding: %v", err))
		return
	}

	// this takes in two values a context and create user parameter
	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating user: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseUsertoUser(user))

}
