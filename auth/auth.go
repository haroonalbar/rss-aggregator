package auth

import (
	"errors"
	"net/http"
	"strings"
)

// Extract ApiKey from a http request header
// Eg : "Authorization": ApiKey <key>
func GetAPIKey(header http.Header) (string, error) {
	val := header.Get("Authorization")
	if val == "" {
		return "", errors.New("no authorization info provided")
	}

	vals := strings.Split(val, " ")
	if len(vals) != 2 {
		return "", errors.New("malformed auth header")
	}

	if vals[0] != "ApiKey" {
		return "", errors.New("malformed fist part of auth header")
	}

	return vals[1], nil
}
