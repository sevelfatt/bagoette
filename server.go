package bagoette

import (
	"net/http"
	"strconv"
)

func (b *BagoetteClient) Serve(opts BagoetteOptions) error {
	if opts.Port != 0  {
		b.opts.Port = opts.Port
	}
	if opts.Host != "" {
		b.opts.Host = opts.Host
	}
	b.registerAllRoutes()
	b.ServeAppearance()
	return http.ListenAndServe(b.opts.Host + ":" + strconv.Itoa(b.opts.Port), b.httpHandler)
}

func (b *BagoetteClient) Close() {
	b.httpClient.CloseIdleConnections()
}