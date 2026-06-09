package bagoette

import (
	"net/http"
)

var Routes []Route

type Route struct {
	name string
	pattern string
	handler http.HandlerFunc
}

type Router struct {
	httpHandler *http.ServeMux
	prefix string
}

func (r *Router) Group(prefix string) *Router {
	newRouter := *r
	newRouter.prefix = r.prefix + prefix
	return &newRouter
}

func (b *BagoetteClient) RegisterRoute(route *Route) {
	b.httpHandler.Handle(route.pattern, route.handler)
}


func (r *Router) AddRoute(route *Route) {
	Routes = append(Routes, *route)
}
