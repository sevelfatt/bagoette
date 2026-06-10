package bagoette

func (r *Router) Group(prefix string) *Router {
	newRouter := *r
	newRouter.prefix = r.prefix + prefix
	return &newRouter
}

func (r *Router) AddRoute(route *Route) {
	*r.routes = append(*r.routes, *route)
}