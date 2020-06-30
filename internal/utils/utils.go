package utils

// BuildRouteURL dd
func BuildRouteURL(route string, params ...string) string {
	result := "/" + route

	for _, param := range params {
		result += "/{" + param + "}"
	}

	return result
}
