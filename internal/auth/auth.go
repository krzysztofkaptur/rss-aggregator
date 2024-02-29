package auth

import (
	"net/http"
)

func GetAPIKey(headers http.Header) (string, error) {
	apiKey := headers.Get("authorization")

	return apiKey, nil
}