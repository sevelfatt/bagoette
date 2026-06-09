package bagoette

import (
	"net/http"
	"strconv"
)

func (b *BagoetteClient) registerAllRoutes() {
	for _, route := range b.routes {
		b.AddRoute(*route)
	}
}

func (b *BagoetteClient) Serve(port int) error {
	b.SetPort(port)
	b.registerAllRoutes()
	return http.ListenAndServe(":" + strconv.Itoa(b.port), b.httpHandler)
}

func (b *BagoetteClient) Close() {
	b.httpClient.CloseIdleConnections()
}