package bagoette

type Route struct {
	name    string
	path    string
	method  string
	pattern string
	handlerFunc HandlerFunc
}

func (r *Router) NewRoute(method string, path string, handlers []HandlerFunc) *Route {
	fn := func(c *Context){
		c.handlers = append(c.handlers, handlers...)
		c.handlers[0](c)
		c.Reset()
	}
	route := &Route{
		method: method,
		path: r.prefix + path,
		pattern: method + " " + r.prefix + path,
		handlerFunc: r.ApplyErrorMiddlewares(fn),
	}	
	r.AddRoute(route)
	return route
}