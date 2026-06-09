package bagoette

import (
	"net/http"

	"github.com/gorilla/mux"
)

type BagoetteClient struct {
	Port string
	HttpClient *http.Client
	MuxRouter *mux.Router
	Router *Router
}

func NewClient() *BagoetteClient {
	return &BagoetteClient{
		HttpClient: &http.Client{},
		MuxRouter: mux.NewRouter(),
		Router: &Router{},
		Port: "8080",
	}
}

func (b *BagoetteClient) SetPort(port string) {
	b.Port = port
}

func (b *BagoetteClient) GetPort() string {
	return b.Port
}


