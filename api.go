package main

import (
	"net/http"
)

func (apiCfg *apiConfig) handlerReadiness(w http.ResponseWriter, r *http.Request) {
	WriteJSON(w, http.StatusOK, struct{}{})
}

func (apiCfg *apiConfig) handlerError(w http.ResponseWriter, r *http.Request) {
	ResponseWithError(w, http.StatusBadRequest, "something went wrong")
}

