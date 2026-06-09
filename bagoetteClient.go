package bagoette

import (
	"net/http"

	"github.com/gorilla/mux"
)

//so this is the main struct of bagoette. 

type BagoetteClient struct {
	Port string // of course this the port that will be used to serve the server
	HttpClient *http.Client // http client will be used to make requests
	MuxRouter *mux.Router //mux router will be used by route to assign new route before passed to http.client
	Routes []*Route // router will be used to route requests
}

func NewClient() *BagoetteClient {
	return &BagoetteClient{
		HttpClient: &http.Client{},
		MuxRouter: mux.NewRouter(),
		Routes: []*Route{},
		Port: "8080",
	}
}

func (b *BagoetteClient) SetPort(port string) {
	b.Port = port
}

func (b *BagoetteClient) GetPort() string {
	return b.Port
}


