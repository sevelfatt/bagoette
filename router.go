package bagoette

import (
	"net/http"
)



type Router struct {
	httpHandler *http.ServeMux
	routes *[]Route
	prefix string
}

func (b *BagoetteClient) NewRouter() *Router {
	return &Router{
		httpHandler: b.httpHandler,
		routes: b.routes,
		prefix: "",
	}
}
