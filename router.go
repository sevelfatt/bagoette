package bagoette

import "github.com/sevelfatt/bagoette/utils"

func (b *BagoetteClient) NewRouter() *Router {
	return &Router{
		httpHandler: b.httpHandler,
		routes:      b.routes,
		middlewares: []HandlerFunc{},
		prefix:      "",
	}
}

func (r *Router) Group(prefix string) *Router {
	newRouter := *r
	newRouter.prefix = r.prefix + prefix
	return &newRouter
}

func (r *Router) Use(handlers ...HandlerFunc) *Router {
	r.middlewares = append(r.middlewares, handlers...)
	return r
}

func (r *Router) NewRoute(method string, path string, handlers []HandlerFunc) *Route {
	pathSegments := utils.GetPathSegment(path)
	paramKeys := utils.GetParamKeys(pathSegments)

	route := &Route{
		Method:   method,
		Path:     r.prefix + path,
		PathSegments: pathSegments,
		ParamKeys: paramKeys,
		Handlers: append(r.middlewares, handlers...),
	}
	r.AddNewRouteToRouter(route)
	return route
}

func (r *Router) AddNewRouteToRouter(route *Route) {
	*r.routes = append(*r.routes, *route)
}
