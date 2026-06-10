package bagoette

import "net/http"

type Route struct {
	name    string
	path    string
	method  string
	pattern string
	handler http.HandlerFunc
}

func (r *Router) NewRoute(method string, path string, handler http.HandlerFunc) *Route {
	route := &Route{
		method: method,
		path: r.prefix + path,
		pattern: method + " " + r.prefix + path,
		handler: r.NotFoundMiddleware(handler),
	}	
	r.AddRoute(route)
	return route
}