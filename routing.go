package bagoette

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	name string
	pattern string
	handler http.HandlerFunc
	muxRoute *mux.Route
}


func (r *Route) SetName(name string) *Route {
	r.name = name
	return r
}

func (b *BagoetteClient) AddRoute(route Route) {
	b.httpHandler.Handle(route.pattern, route.handler)
	b.routes = append(b.routes, &route)
}
