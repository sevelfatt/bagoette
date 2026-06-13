package bagoette

import (
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

	b.httpServer.Handler = b.httpHandler
	b.httpServer.Addr = b.opts.Host + ":" + strconv.Itoa(b.opts.Port)

	return b.httpServer.ListenAndServe()
}

func (b *BagoetteClient) Close() error {
	return b.httpServer.Close()
}