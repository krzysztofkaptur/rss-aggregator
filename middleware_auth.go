package main

import (
	"fmt"
	"net/http"

	"github.com/krzysztofkaptur/rss-aggregator/internal/auth"
	"github.com/krzysztofkaptur/rss-aggregator/internal/database"
)

type authHandler func(http.ResponseWriter, *http.Request, database.User)

func (apiCfg *apiConfig) middlewareAuth(handler authHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		handler(w, r, user)
	}
}