package bagoette

import (
	"net/http"
)

func (b *BagoetteClient) registerAllRoutes() {
	for _, route := range b.Routes {
		b.AddRoute(*route)
	}
}

func (b *BagoetteClient) Serve(port string) error {
	b.SetPort(port)
	b.registerAllRoutes()
	return http.ListenAndServe(":" + b.Port, b.MuxRouter)
}

func (b *BagoetteClient) Close() {
	b.HttpClient.CloseIdleConnections()
}