package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/google/uuid"
	"github.com/haroonalbar/rss-aggregater/internal/database"
)

func (apiCfg *apiConfig) handlerCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	params := Parameters{}
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error decoding: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.Url,
		UserID:    user.ID,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error creating feed: %v", err))
		return
	}

	respondWithJSON(w, 201, databaseFeedtoFeed(feed))

}

func (apiCfg *apiConfig) handlerUpdateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type Parameters struct {
		Name string `json:"name"`
		Url  string `json:"url"`
	}
	feedIDStr := chi.URLParam(r, "feedID")
	feedID, err := uuid.Parse(feedIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing feed id from url: %v", err))
		return
	}
	params := Parameters{}
	decoder := json.NewDecoder(r.Body)
	err = decoder.Decode(&params)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error decoding: %v", err))
		return
	}

	feed, err := apiCfg.DB.UpdateFeed(r.Context(), database.UpdateFeedParams{
		ID:     feedID,
		UserID: user.ID,
		Name:   params.Name,
		Url:    params.Url,
	})
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error updating feed: %v", err))
		return
	}

	respondWithJSON(w, 200, databaseFeedtoFeed(feed))
}

func (apiCfg *apiConfig) handlerGetFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeeds(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get feeds: %v", err))
		return
	}
	respondWithJSON(w, 200, databaseFeedstoFeeds(feeds))
}

func (apiCfg *apiConfig) handlerGetNextFeedToFetch(w http.ResponseWriter, r *http.Request) {
	nextFeed, err := apiCfg.DB.GetNextFeedToFetch(r.Context())
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get next feed to fetch: %v", err))
		return
	}
	respondWithJSON(w, 200, nextFeed)
}

func (apiCfg *apiConfig) handlerMarkFeedAsFetched(w http.ResponseWriter, r *http.Request) {
	feedIDStr := chi.URLParam(r, "feedID")
	feedID, err := uuid.Parse(feedIDStr)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Error parsing feed id: %v", err))
		return
	}
	nextFeed, err := apiCfg.DB.MarkFeedAsFetched(r.Context(), feedID)
	if err != nil {
		respondWithError(w, 400, fmt.Sprintf("Couldn't get next feed to fetch: %v", err))
		return
	}
	respondWithJSON(w, 200, nextFeed)
}
