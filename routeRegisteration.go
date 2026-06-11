package bagoette

import (
	"net/http"
)

func (b *BagoetteClient) registerAllRoutes() {
	for _, route := range *b.routes {
		go b.registerRoute(&route)
	}
}

func (b *BagoetteClient) registerRoute(route *Route) {
	fn := func(w http.ResponseWriter, r *http.Request){
		b.context.w = w
		b.context.r = r
		b.context.currentRoute = route
		route.Handlers[0](b.context)
	}

	b.httpHandler.Handle(route.pattern, http.HandlerFunc(fn))
}
