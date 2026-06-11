package utils

import (
	"net/http"
	"strings"

)

func MatchRoute(routePathSegments []string, requestPathSegments []string) bool {
	if len(routePathSegments) != len(requestPathSegments) {
		return false
	}

	for i := range routePathSegments {
		routePart := routePathSegments[i]
		
		if strings.HasPrefix(routePart, "{") && strings.HasSuffix(routePart, "}") {
			continue
		}

		if routePart != requestPathSegments[i] {
			return false
		}
	}
	return true
}

func MatchRouteMethod(method string, request *http.Request) bool {
	return method == request.Method
}

func GetPathSegment(path string) []string {
	return strings.Split(strings.Trim(path, "/"), "/")
}

func GetParamKeys(routePathSegments []string) []string {
	var paramKeys []string
	for _, routePart := range routePathSegments {
		if strings.HasPrefix(routePart, "{") && strings.HasSuffix(routePart, "}") {
			paramKeys = append(paramKeys, routePart[1:len(routePart)-1])
		}
	}
	return paramKeys
}