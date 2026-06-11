package bagoette

import (
	"net/http"
)

//NewClient: create a new BagoetteClient
func NewClient() *BagoetteClient {
	return &BagoetteClient{
		httpClient: http.DefaultClient,
		httpHandler: http.NewServeMux(),
		routes: &[]Route{},
		router: &Router{},
		context: &Context{},
		port: 8080,
	}
}

//SetPort: set the port of the server
func (b *BagoetteClient) SetPort(port int) {
	b.port = port
}

//GetPort: get the port of the server
func (b *BagoetteClient) GetPort() int {
	return b.port
}


