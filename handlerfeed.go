package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/KatragaddaLokesh/RSSagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type param struct {
		Name string `json:"name"`
		URL  string `json: "url"`
	}
	decode := json.NewDecoder(r.Body)

	params := param{}
	err := decode.Decode(&params)

	if err != nil {
		responwithError(w, 400, fmt.Sprintf("Error parasing Json: %v", err))
		return
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
		Url:       params.URL,
		UserID:    user.ID,
	})
	if err != nil {
		responwithError(w, 400, fmt.Sprintf("Could't Create Feed: %v", err))
		return
	}

	responWithJSON(w, 201, databaseFeedToFeed(feed))

}

func (apiCfg *apiConfig) handleGetFeed(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.GetFeed(r.Context())
	if err != nil {
		responwithError(w, 400, fmt.Sprintf("Could't Get Feed: %v", err))
		return
	}

	responWithJSON(w, 201, databaseFeedToFeeds(feeds))

}
