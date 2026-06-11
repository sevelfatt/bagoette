package bagoette

import (
	"fmt"
	"net/http"
	"time"
)

func (b *BagoetteClient) registerAllRoutes() {
	start := time.Now()
	for _, route := range *b.routes {
		go b.registerRoute(&route)
	}
	fmt.Println("Registered all routes in", time.Since(start))
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
