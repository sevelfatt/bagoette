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
		context.w = w
		context.r = r
		context.currentRoute = route
		route.Handlers[0](context)
	}

	b.httpHandler.Handle(route.Method + " " + route.Path, http.HandlerFunc(fn))
}
