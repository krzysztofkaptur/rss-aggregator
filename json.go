package main

import (
	"encoding/json"
	"log"
	"net/http"
)

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

func ResponseWithError(w http.ResponseWriter, statusCode int, msg string) {
	if statusCode >= 500 {
		log.Println("responding with 5XX error:", msg)
	}

	type errResponse struct {
		Error string `json:"error"`
	}

	WriteJSON(w, statusCode, errResponse{Error: msg})
}