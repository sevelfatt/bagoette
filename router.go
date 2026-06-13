package bagoette

import (
	"github.com/sevelfatt/bagoette/utils"
)

//Router struct: work as the router of the server
type RouteGroup struct {
	routes *[]Route
	middlewares []HandlerFunc
	prefix string
}

//Route struct: Define individual route
type Route struct {
	path    string
	method  string

	pathSegments []string
	paramKeys []string

	handlers []HandlerFunc
}

func (b *BagoetteClient) NewRouter() *RouteGroup {
	return &RouteGroup{
		routes:      b.routes,
		middlewares: []HandlerFunc{b.NotFoundMiddleware, b.MethodNotAllowedMiddleware, b.InternalServerErrorMiddleware},
		prefix:      "",
	}
}

func (r *RouteGroup) Group(prefix string) *RouteGroup {
	newRouter := *r
	newRouter.prefix = r.prefix + prefix
	return &newRouter
}

func (r *RouteGroup) Use(handlers ...HandlerFunc) *RouteGroup {
	r.middlewares = append(r.middlewares, handlers...)
	return r
}

func (r *RouteGroup) NewRoute(method string, path string, handlers []HandlerFunc) {
	fullPath := r.prefix + path
	
	pathSegments := utils.GetPathSegment(fullPath)
	paramKeys := utils.GetParamKeys(pathSegments)

	route := &Route{
		method:   method,
		path:     fullPath,
		pathSegments: pathSegments,
		paramKeys: paramKeys,
		handlers: append(r.middlewares, handlers...),
	}
	r.AddNewRouteToBagoetteClient(route)
}

func (r *RouteGroup) AddNewRouteToBagoetteClient(route *Route) {
	*r.routes = append(*r.routes, *route)
}
