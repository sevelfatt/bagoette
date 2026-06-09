package bagoette

import (
	"net/http"

	"github.com/gorilla/mux"
)

type BagoetteClient struct {
	Port string
	Host string
	HttpClient *http.Client
	MuxRouter *mux.Router
}

func NewClient() *BagoetteClient {
	return &BagoetteClient{
		HttpClient: &http.Client{},
		MuxRouter: mux.NewRouter(),
		Host: "localhost",
		Port: "8080",
	}
}

func (b *BagoetteClient) SetPort(port string) {
	b.Port = port
}

func (b *BagoetteClient) SetHost(host string) {
	b.Host = host
}

func (b *BagoetteClient) GetPort() string {
	return b.Port
}

func (b *BagoetteClient) GetHost() string {
	return b.Host
}

