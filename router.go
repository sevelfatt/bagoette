package bagoette

import (
	"net/http"
)

type Route struct {
	name string
	path string
	method string
	pattern string
	handler http.HandlerFunc
}

type Router struct {
	httpHandler *http.ServeMux
	routes *[]Route
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
	*r.routes = append(*r.routes, *route)
}
