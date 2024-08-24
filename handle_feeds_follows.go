package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/KatragaddaLokesh/RSSagg/internal/database"
	"github.com/go-chi/chi"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type param struct {
		FeedID uuid.UUID `json:"feed_id"`
	}
	decode := json.NewDecoder(r.Body)

	params := param{}
	err := decode.Decode(&params)

	if err != nil {
		responwithError(w, 400, fmt.Sprintf("Error parasing Json: %v", err))
		return
	}

	feedfollow, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		UserID:    user.ID,
		FeedID:    params.FeedID,
	})
	if err != nil {
		responwithError(w, 400, fmt.Sprintf("Could't Create FeedFollow: %v", err))
		return
	}

	responWithJSON(w, 201, databaseFeedFollowersToFeedFollowers(feedfollow))

}

func (apiCfg *apiConfig) handleGetFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	feedfollow, err := apiCfg.DB.GetFeedFollows(r.Context(), user.ID)
	if err != nil {
		responwithError(w, 400, fmt.Sprintf("Could't get FeedFollow: %v", err))
		return
	}

	responWithJSON(w, 201, databaseFeedFollowsToFeedsFollows(feedfollow))

}

func (apiCfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	FeedsFollowIDstr := chi.URLParam(r, "feedFollowID")

	FeedsFollowID, err := uuid.Parse(FeedsFollowIDstr)
	if err != nil {
		responwithError(w, 400, fmt.Sprintf("Could't get Parse FeedFollow: %v", err))
		return
	}

	err = apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{
		ID:     FeedsFollowID,
		UserID: user.ID,
	})
	if err != nil {
		responwithError(w, 400, fmt.Sprintf("Could't get Delete FeedFollow: %v", err))
		return
	}

	responWithJSON(w, 200, struct{}{})
}
