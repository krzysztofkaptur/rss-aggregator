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

	// examples
	router.HandleFunc("GET /api/v1/healthz", apiCfg.handlerReadiness)
	router.HandleFunc("GET /api/v1/error", apiCfg.handlerError)

	// users
	router.HandleFunc("POST /api/v1/users", apiCfg.handleCreateUser)
	router.HandleFunc("GET /api/v1/users", apiCfg.middlewareAuth(apiCfg.handleFetchUser))

	// feeds
	router.HandleFunc("GET /api/v1/feeds", apiCfg.handleFetchFeeds)
	router.HandleFunc("POST /api/v1/feeds", apiCfg.middlewareAuth(apiCfg.handleCreateFeed))
	
	// feed follows
	router.HandleFunc("GET /api/v1/feed-follow", apiCfg.middlewareAuth(apiCfg.handleFetchFeedFollows))
	router.HandleFunc("POST /api/v1/feed-follow", apiCfg.middlewareAuth(apiCfg.handleCreateFeedFollow))
	router.HandleFunc("DELETE /api/v1/feed-follow/{feedId}", apiCfg.middlewareAuth(apiCfg.handleDeleteFeedFollow))


	server := &http.Server{
		Handler: router,
		Addr: ":" + os.Getenv("PORT"),
	}

	serverErr := server.ListenAndServe()
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
