package main

import "net/http"

func handlerReadiness(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, struct{}{})
}

func handlerError(w http.ResponseWriter, r *http.Request) {
	ResponseWithError(w, http.StatusBadRequest, "something went wrong")
}