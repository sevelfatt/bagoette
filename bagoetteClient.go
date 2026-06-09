package bagoette

import (
	"net/http"
)

//so this is the main struct of bagoette. 

type BagoetteClient struct {
	port int // of course this the port that will be used to serve the server
	httpClient *http.Client // http client will be used to make requests
	httpHandler *http.ServeMux
	router *Router // router will be used to route requests
}

func NewClient() *BagoetteClient {
	return &BagoetteClient{
		httpClient: http.DefaultClient,
		httpHandler: http.NewServeMux(),
		router: &Router{},
		port: 8080,
	}
}

func (b *BagoetteClient) NewRouter() *Router {
	return &Router{
		httpHandler: b.httpHandler,
	}
}

func (b *BagoetteClient) SetPort(port int) {
	b.port = port
}

func (b *BagoetteClient) GetPort() int {
	return b.port
}


