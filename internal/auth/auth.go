package auth

import (
	"errors"
	"net/http"
	"strings"
)

var ErrNoAuthHeaderIncluded error = errors.New("no authorization header included")

// To get authentication key from request headers
func GetApiKey(header http.Header) (string, error) {
	authHeader := header.Get("Authorization")
	if authHeader == "" {
		return "", ErrNoAuthHeaderIncluded
	}
	splitAuthHeader := strings.Split(authHeader, " ")
	if len(splitAuthHeader) < 2 || splitAuthHeader[0] != "ApiKey" {
		return "", ErrNoAuthHeaderIncluded
	}
	return splitAuthHeader[1], nil
}
