package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)


func main() {
	InitEnv()
	
	store, err := InitDB()
	if err != nil {
		log.Fatal(err)
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/healthz", handlerReadiness)
	router.HandleFunc("GET /api/v1/error", handlerError)

	router.HandleFunc("GET /api/v1/users", func (w http.ResponseWriter, r *http.Request) {
		users, err := store.FetchUsers(r.Context())
		if err != nil {
			log.Fatal(err)
		}

		WriteJSON(w, http.StatusOK, users)
	})

	server := &http.Server{
		Handler: router,
		Addr: ":" + os.Getenv("PORT"),
	}

	serverErr := server.ListenAndServe()
	if serverErr != nil {
		log.Fatal(serverErr)
	}
}
