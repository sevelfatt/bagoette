package bagoette

import (
	"net/http"
)

func (b *BagoetteClient) registerAllRoutes() {
	for _, route := range *b.routes {
		b.registerRoute(&route)
	}
}

func (b *BagoetteClient) registerRoute(route *Route) {
	fn := func(w http.ResponseWriter, r *http.Request){
		context := NewContext()
		context.writer = w
		context.request = r
		context.currentRoute = route
		route.handlers[0](context)
	}

	b.httpHandler.Handle(route.method + " " + route.path, http.HandlerFunc(fn))
}
