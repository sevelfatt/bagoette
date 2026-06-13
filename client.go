package bagoette

import (
	"net/http"

)

//NewClient: create a new BagoetteClient
func NewClient() *BagoetteClient {
	return &BagoetteClient{
		httpServer: &http.Server{},
		httpHandler: http.NewServeMux(),
		routes: &[]Route{},
		opts: &BagoetteOptions{
			Port: 8080,
			Host: "localhost",
		},
	}
}

func (b *BagoetteClient) GetOpts() BagoetteOptions {
	return *b.opts
}

func (b *BagoetteClient) GetRoutes() []Route {
	return *b.routes
}

func (b *BagoetteClient) GetHTTPHandler() *http.ServeMux {
	return b.httpHandler
}
