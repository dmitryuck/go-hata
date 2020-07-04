package utils

import (
	"errors"
	"net/http"
)

// BuildRouteURL dd
func BuildRouteURL(route string, params ...string) string {
	result := "/" + route

	for _, param := range params {
		result += "/{" + param + "}"
	}

	return result
}

// GetQueryParam get param from query string
func GetQueryParam(r *http.Request, key string) (string, error) {
	keys, ok := r.URL.Query()[key]

	if !ok || len(keys[0]) < 1 {
		return "", errors.New("Url Param " + key + " is missing")
	}

	return keys[0], nil
}
