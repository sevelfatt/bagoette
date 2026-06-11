package bagoette

import (
	"net/http"
	"strconv"
)

func (b *BagoetteClient) Serve(port int) error {
	b.SetPort(port)
	b.registerAllRoutes()
	b.ServeAppearance()
	return http.ListenAndServe(":" + strconv.Itoa(b.port), b.httpHandler)
}

func (b *BagoetteClient) Close() {
	b.httpClient.CloseIdleConnections()
}