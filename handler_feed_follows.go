package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/krzysztofkaptur/rss-aggregator/internal/database"
)

func (apiCfg *apiConfig) handleCreateFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	type parameters struct {
		FeedId int32 `json:"feed_id"`
	} 

	params := parameters{}

	err := json.NewDecoder(r.Body).Decode(&params)
	if err != nil {
		// TODO: handle errors better
		fmt.Println(err)
	}

	feed, err := apiCfg.DB.CreateFeedFollow(r.Context(), database.CreateFeedFollowParams{ UserID: user.ID, FeedID: params.FeedId })	
	if err != nil {
		// TODO: handle errors better
		fmt.Println(err)
	}

	WriteJSON(w, http.StatusCreated, feed)
}

func (apiCfg *apiConfig) handleFetchFeedFollows(w http.ResponseWriter, r *http.Request, user database.User) {
	feeds, err := apiCfg.DB.FetchFeedFollows(r.Context(), user.ID)
	if err != nil {
		// TODO: handle errors better
		fmt.Println(err)
	}

	WriteJSON(w, http.StatusOK, feeds)
}

func (apiCfg *apiConfig) handleDeleteFeedFollow(w http.ResponseWriter, r *http.Request, user database.User) {
	strFeedId := r.PathValue("feedId")
	feedId, err := strconv.Atoi(strFeedId)
	if err != nil {
		// TODO: handle errors better
		fmt.Println(err)
	}

	feed, err := apiCfg.DB.DeleteFeedFollows(r.Context(), database.DeleteFeedFollowsParams{UserID: user.ID, FeedID: int32(feedId)})
	if err != nil {
		// TODO: handle errors better
		fmt.Println(err)
	}

	WriteJSON(w, http.StatusOK, feed)
}