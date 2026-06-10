package bagoette

import "net/http"

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

func (r *Router) Get(path string, handler http.HandlerFunc) *Route {
	return r.NewRoute("GET", path, handler)
}

func (r *Router) Post(path string, handler http.HandlerFunc) *Route {
	return r.NewRoute("POST", path, handler)
}

func (r *Router) Put(path string, handler http.HandlerFunc) *Route {
	return r.NewRoute("PUT", path, handler)
}

func (r *Router) Delete(path string, handler http.HandlerFunc) *Route {
	return r.NewRoute("DELETE", path, handler)
}

func (r *Router) Patch(path string, handler http.HandlerFunc) *Route {
	return r.NewRoute("PATCH", path, handler)
}

func (r *Router) Head(path string, handler http.HandlerFunc) *Route {
	return r.NewRoute("HEAD", path, handler)
}

func (r *Router) Options(path string, handler http.HandlerFunc) *Route {
	return r.NewRoute("OPTIONS", path, handler)
}

func (r *Router) Connect(path string, handler http.HandlerFunc) *Route {
	return r.NewRoute("CONNECT", path, handler)
}

func (r *Router) Trace(path string, handler http.HandlerFunc) *Route {
	return r.NewRoute("TRACE", path, handler)
}