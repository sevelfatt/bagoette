package bagoette

import (
	"strconv"
)

func (b *BagoetteClient) Serve() error {
	b.registerAllRoutes()
	b.ServeAppearance()

	b.httpServer.Handler = b.httpHandler
	b.httpServer.Addr = ":" + strconv.Itoa(b.Opts.Port)

	return b.httpServer.ListenAndServe()
}

func (b *BagoetteClient) Close() error {
	return b.httpServer.Close()
}
