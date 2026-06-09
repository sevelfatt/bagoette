package bagoette

import "net/http"

func (b *BagoetteClient) Get(path string, handler http.HandlerFunc) *Route {
	pattern := "GET " + path
	route := &Route{
		pattern: pattern,
		handler: handler,
	}
	b.AddRoute(route)
	return route
}