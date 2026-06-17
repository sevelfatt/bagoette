package bagoette

import (
	"errors"
	"net/http"
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
	if port < 0 || port > 65535 {
		Logger.Println("Error: Invalid port, using default port 8080")
		return b
		
	}
	b.Opts.Port = port
	return b
}

//MaxUploadSize: set the max upload size of the server
func (b *BagoetteClient) MaxUploadSize(size int64) (*BagoetteClient, error) {
	if size < 0 {
		return nil, errors.New("Invalid max upload size")
	}
	b.Opts.MaxUploadSize = size
	return b, nil
}

func (b *BagoetteClient) GetRoutes() []Route {
	return *b.routes
}

func (b *BagoetteClient) GetHTTPHandler() *http.ServeMux {
	return b.httpHandler
}
