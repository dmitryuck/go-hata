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
	values, ok := r.URL.Query()[key]

	if !ok || len(values[0]) < 1 {
		return "", errors.New("Url Param " + key + " is missing")
	}

	return values[0], nil
}

// GetQueryParams get query params
func GetQueryParams(r *http.Request, keys ...string) (map[string]string, error) {
	result := make(map[string]string)

	for _, key := range keys {
		values, ok := r.URL.Query()[key]

		if !ok || len(values[0]) < 1 {
			return nil, errors.New("Url Param " + key + " is missing")
		}

		result[key] = values[0]
	}

	return result, nil
}
