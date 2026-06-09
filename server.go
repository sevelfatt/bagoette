package bagoette

import (
	"net/http"
)

func (b *BagoetteClient) ListenAndServe() error {
	return http.ListenAndServe(b.Host + ":" + b.Port, b.MuxRouter)
}

func (b *BagoetteClient) Close() {
	b.HttpClient.CloseIdleConnections()
}