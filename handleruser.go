package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/KatragaddaLokesh/RSSagg/internal/database"
	"github.com/google/uuid"
)

func (apiCfg *apiConfig) handleUser(w http.ResponseWriter, r *http.Request) {
	type param struct {
		Name string `json:"name"`
	}
	decode := json.NewDecoder(r.Body)

	params := param{}
	err := decode.Decode(&params)

	if err != nil {
		responwithError(w, 400, fmt.Sprintf("Error parasing Json: %v", err))
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
		Name:      params.Name,
	})
	if err != nil {
		responwithError(w, 400, fmt.Sprintf("Could't Create User: %v", err))
		return
	}

	responWithJSON(w, 201, databaseUserToUser(user))

}

func (apiCfg *apiConfig) handleGetUser(w http.ResponseWriter, r *http.Request, user database.User) {
	responWithJSON(w, 201, databaseUserToUser(user))
}

func (apiCfg *apiConfig) handleGetPostForUsers(w http.ResponseWriter, r *http.Request, user database.User) {
	post, err := apiCfg.DB.GetPostForUser(r.Context(), database.GetPostForUserParams{
		UserID: user.ID,
		Limit:  10,
	})
	if err != nil {
		responwithError(w, 400, fmt.Sprintf("Could't Get Posts: %v", err))
		return
	}

	responWithJSON(w, 200, databasePostsToPosts(post))
}
