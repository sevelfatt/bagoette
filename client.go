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
		opts: &BagoetteOptions{
			Port: 8080,
			Host: "localhost",
		},
	}
}


