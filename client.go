package bagoette

import (
	"net/http"
	"os"
	"strconv"
)

//Bagoette main struct: work as the core of the library
//and provide all the features like router, middleware, context, etc
type BagoetteClient struct {
	Opts *Options

	httpServer *http.Server
	httpHandler *http.ServeMux

	routes *[]Route
}

//NewClient: create a new BagoetteClient
func NewClient() *BagoetteClient {
	return &BagoetteClient{
		httpServer: &http.Server{},
		httpHandler: http.NewServeMux(),
		routes: &[]Route{},
		Opts: NewOptions(),
	}
}

//Port: set the port of the server
func (b *BagoetteClient) Port(port int) *BagoetteClient {
	b.Opts.Port = port
	return b
}

//MaxUploadSize: set the max upload size of the server
func (b *BagoetteClient) MaxUploadSize(size int64) *BagoetteClient {
	b.Opts.MaxUploadSize = size
	return b
}

//getDefaultPort: get the default port from the environment variable PORT
func getDefaultPort() int {
	port := os.Getenv("PORT")
	if port == "" {
		logger.Println("Warning: No port in environment variable, using default port 8080")
		return 8080
	}
	parsedPort, err := strconv.Atoi(port)
	if err != nil {
		logger.Println("Warning: Invalid port in environment variable, using default port 8080")
		return 8080
	}
	return parsedPort
}

func (b *BagoetteClient) GetRoutes() []Route {
	return *b.routes
}

func (b *BagoetteClient) GetHTTPHandler() *http.ServeMux {
	return b.httpHandler
}
