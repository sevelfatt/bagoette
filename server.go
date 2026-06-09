package bagoette

import (
	"net/http"
)

func (b *BagoetteClient) ListenAndServe(port string) error {
	b.SetPort(port)
	return http.ListenAndServe(":" + b.Port, b.MuxRouter)
}

func (b *BagoetteClient) Close() {
	b.HttpClient.CloseIdleConnections()
}