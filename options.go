package bagoette

import (
	"os"
	"strconv"
)

type Options struct {
	Port          int
	MaxUploadSize int64
}

func NewOptions() *Options {
	return &Options{
		Port:          getDefaultPort(),
		MaxUploadSize: 10 * 1024 * 1024, // 10MB
	}
}

// getDefaultPort: get the default port from the environment variable PORT
func getDefaultPort() int {
	port := os.Getenv("PORT")
	if port == "" {
		Logger.Println("Warning: No port in environment variable, using default port 8080")
		return 8080
	}
	parsedPort, err := strconv.Atoi(port)
	if parsedPort < 0 || parsedPort > 65535 {
		Logger.Printf("Warning: Invalid port %d in environment variable, using default port 8080", parsedPort)
		return 8080
	}
	if err != nil {
		Logger.Printf("Warning: Invalid port %s in environment variable, using default port 8080", port)
		return 8080
	}
	return parsedPort
}
