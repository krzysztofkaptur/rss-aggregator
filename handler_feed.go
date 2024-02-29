package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/krzysztofkaptur/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handleCreateFeed(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		Name string `json:"name"`
		Url string `json:"url"`
	}

	params := &parameters{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		// todo: handle error
		fmt.Println(err)
	}

	feed, err := apiCfg.DB.CreateFeed(r.Context(), database.CreateFeedParams{Name: params.Name, Url: params.Url, UserID: user.ID})
	if err != nil {
		// todo: handle error
		fmt.Println(err)
	}

	WriteJSON(w, http.StatusCreated, feed)
}

func (apiCfg *apiConfig) handleFetchFeeds(w http.ResponseWriter, r *http.Request) {
	feeds, err := apiCfg.DB.FetchFeeds(r.Context())
	if err != nil {
		// todo: handle error
		fmt.Println(err)
	}

	WriteJSON(w, http.StatusOK, feeds)
}