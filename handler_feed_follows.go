package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/haroonalbar/rss-aggregater/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameters struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	params := Parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error decoding: %v", err))
		return
	}

	feedFollow, err := apiCfg.DB.CreateFeedFollows(r.Context(), database.CreateFeedFollowsParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:       params.FeedID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating feed_follows: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedFollowstoFeedFollows(feedFollow))

}

