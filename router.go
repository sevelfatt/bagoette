package bagoette

func (b *BagoetteClient) NewRouter() *Router {
	return &Router{
		httpHandler: b.httpHandler,
		routes: b.routes,
		context: b.context,
		middlewares: []HandlerFunc{FirstHandlerMiddleware},
		prefix: "",
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
	route := &Route{
		method: method,
		path: r.prefix + path,
		pattern: method + " " + r.prefix + path,
		Handlers: append(r.middlewares, handlers...),
	}	
	r.AddNewRouteToRouter(route)
	return route
}

func (r *Router) AddNewRouteToRouter(route *Route) {
	*r.routes = append(*r.routes, *route)
}
