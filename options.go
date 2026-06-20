package bagoette

import (
	"os"
	"strconv"
)

type Options struct {
	Port          int
	MaxUploadSize int64
	UseCors       bool
	Cors          *Cors
}

func NewOptions() *Options {
	return &Options{
		Port:          getDefaultPort(),
		MaxUploadSize: 10 * 1024 * 1024, // 10MB
		UseCors:       false,
		Cors:          nil,
	}
}

func (b *BagoetteClient) SetOptions(opts Options) *BagoetteClient {
	if opts.Port != 0 {
		b.Opts.Port = opts.Port
	}
	if opts.MaxUploadSize != 0 {
		b.Opts.MaxUploadSize = opts.MaxUploadSize
	}
	if opts.UseCors != false {
		b.Opts.UseCors = opts.UseCors
	}
	if opts.Cors != nil {
		b.Opts.Cors = opts.Cors
	}
	return b
}

func (b *BagoetteClient) SetCors(cors *Cors) *BagoetteClient {
	b.Opts.UseCors = true
	if cors != nil {
		b.Opts.Cors = cors
	} else {
		b.Opts.Cors = NewCors()
	}
	return b
}

func (b *BagoetteClient) DisableCors() *BagoetteClient {
	b.Opts.UseCors = false
	b.Opts.Cors = nil
	return b
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
