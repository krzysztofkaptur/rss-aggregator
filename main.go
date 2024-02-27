package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := http.NewServeMux()

	router.HandleFunc("GET /api/v1/healthz", handlerReadiness)
	router.HandleFunc("GET /api/v1/error", handlerError)

	server := &http.Server{
		Handler: router,
		Addr: ":" + os.Getenv("PORT"),
	}

	err = server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func WriteJSON(w http.ResponseWriter, statusCode int, payload any) {
	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("failed to marshal JSON response: %v", payload)
		w.WriteHeader(500)
		return
	}
	
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	w.Write(data)
}

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, struct{}{})
}

func ResponseWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode >= 500 {
		log.Println("responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	WriteJSON(w, statusCode, errResponse{Error: msg})
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	ResponseWithError(w, http.StatusBadRequest, "something went wrong")
}