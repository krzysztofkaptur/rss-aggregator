package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type createUserParams struct {
	Email string `json:"email"`
} 

func (apiCfg *apiConfig) handleFetchUsers(w http.ResponseWriter, r *http.Request) {
	users, err := apiCfg.DB.FetchUsers(r.Context())

	if err != nil {
		log.Fatal(err)
	}

	WriteJSON(w, http.StatusOK, users)
}

func (apiCfg *apiConfig) handleCreateUser(w http.ResponseWriter, r *http.Request) {
	createUserPar := createUserParams{}
	err := json.NewDecoder(r.Body).Decode(&createUserPar)
	if err != nil {
		fmt.Println("something wrong with json")
		return
	}

	user, err := apiCfg.DB.CreateUser(r.Context(), createUserPar.Email)
	if err != nil {
		fmt.Println("something wrong with user creation")
		return
	}

	WriteJSON(w, http.StatusOK, user)
}