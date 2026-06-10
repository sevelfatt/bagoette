package bagoette

func (r *Router) Get(path string, handlers ...HandlerFunc) *Route {
	return r.NewRoute("GET", path, handlers)
}

func (r *Router) Post(path string, handlers ...HandlerFunc) *Route {
	return r.NewRoute("POST", path, handlers)
}

func (r *Router) Put(path string, handlers ...HandlerFunc) *Route {
	return r.NewRoute("PUT", path, handlers)
}

func (r *Router) Delete(path string, handlers ...HandlerFunc) *Route {
	return r.NewRoute("DELETE", path, handlers)
}

func (r *Router) Patch(path string, handlers ...HandlerFunc) *Route {
	return r.NewRoute("PATCH", path, handlers)
}

func (r *Router) Head(path string, handlers ...HandlerFunc) *Route {
	return r.NewRoute("HEAD", path, handlers)
}

func (r *Router) Options(path string, handlers ...HandlerFunc) *Route {
	return r.NewRoute("OPTIONS", path, handlers)
}

func (r *Router) Connect(path string, handlers ...HandlerFunc) *Route {
	return r.NewRoute("CONNECT", path, handlers)
}

func (r *Router) Trace(path string, handlers ...HandlerFunc) *Route {
	return r.NewRoute("TRACE", path, handlers)
}