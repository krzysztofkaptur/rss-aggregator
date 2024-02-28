package auth

import (
	"fmt"
	"net/http"
)

// TODO: change it into middleware
func GetAPIKey(headers http.Header) (string, error) {
	fmt.Println(headers.Get("authorization"))

	return "something", nil
}