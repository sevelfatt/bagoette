package bagoette

import (
	"net/http"
)

type Route struct {
	name string
	pattern string
	handler http.HandlerFunc
}


func (r *Route) SetName(name string) *Route {
	r.name = name
	return r
}

func (b *BagoetteClient) RegisterRoute(route *Route) {
	b.httpHandler.Handle(route.pattern, route.handler)
}

func (b *BagoetteClient) GetRoutes() []*Route {
	return b.routes
}

func (b *BagoetteClient) AddRoute(route *Route) {
	b.routes = append(b.routes, route)
}
