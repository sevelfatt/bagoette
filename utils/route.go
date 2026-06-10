package utils

import "strings"

func MatchRoute(pattern string, path string) bool {
	patternParts := strings.Split(strings.Trim(pattern, "/"), "/")
	pathParts := strings.Split(strings.Trim(path, "/"), "/")

	if len(patternParts) != len(pathParts) {
		return false
	}

	for i := range patternParts {
		part := patternParts[i]
		
		if strings.HasPrefix(part, "{") && strings.HasSuffix(part, "}") {
			continue
		}

		if part != pathParts[i] {
			return false
		}
	}
	return true
}