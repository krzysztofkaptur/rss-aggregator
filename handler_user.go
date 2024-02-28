package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/krzysztofkaptur/rss-aggregator/internal/auth"
)

type createUserParams struct {
	Email string `json:"email"`
} 

func (apiCfg *apiConfig) handleFetchUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.FetchUsers(r.Context())

	if err != nil {
		// TODO: handle errors with ResponseWithError
		log.Fatal(err)
	}

	WriteJSON(w, http.StatusOK, users)
}

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	createUserPar := createUserParams{}
	err := json.NewDecoder(r.Body).Decode(&createUserPar)
	if err != nil {
		// TODO: handle errors with ResponseWithError
		fmt.Println("something wrong with json")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), createUserPar.Email)
	if err != nil {
		// TODO: handle errors with ResponseWithError
		fmt.Println("something wrong with user creation")
		return
	}

	WriteJSON(w, http.StatusCreated, user)
}

func (apiCfg *apiConfig) handleFetchUser(w http.ResponseWriter, r *http.Request) {
	apiKey, err := auth.GetAPIKey(r.Header)
	if err != nil {
		// TODO: handle errors with ResponseWithError
		fmt.Println("wrong api key")
		return
	}

	user, err := apiCfg.DB.GetUserByAPIKey(r.Context(), apiKey)
	if err != nil {
		// TODO: handle errors with ResponseWithError 
		fmt.Println("user doesn't exist")
		return
	}

	WriteJSON(w, http.StatusOK, user)
}