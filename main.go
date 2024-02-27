package main

import (
	"log"
	"net/http"
	"os"

	"github.com/krzysztofkaptur/rss-aggregator/internal/database"
	_ "github.com/lib/pq"
)

type apiConfig struct {
	DB *database.Queries
}

func main() {
	InitEnv()
	
	store, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	apiCfg := apiConfig{
		DB: store,
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/healthz", apiCfg.handlerReadiness)
	router.HandleFunc("GET /api/v1/error", apiCfg.handlerError)

	router.HandleFunc("GET /api/v1/users", apiCfg.handleFetchUsers)
	router.HandleFunc("POST /api/v1/users", apiCfg.handleCreateUser)

	server := &http.Server{
		Handler: router,
		Addr: ":" + os.Getenv("PORT"),
	}

	serverErr := server.ListenAndServe()
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
