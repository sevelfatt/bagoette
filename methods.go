package bagoette

func (r *RouteGroup) Get(path string, handlers ...HandlerFunc) {
	r.NewRoute("GET", path, handlers)
}

func (r *RouteGroup) Post(path string, handlers ...HandlerFunc) {
	r.NewRoute("POST", path, handlers)
}

func (r *RouteGroup) Put(path string, handlers ...HandlerFunc) {
	r.NewRoute("PUT", path, handlers)
}

func (r *RouteGroup) Delete(path string, handlers ...HandlerFunc) {
	r.NewRoute("DELETE", path, handlers)
}

func (r *RouteGroup) Patch(path string, handlers ...HandlerFunc) {
	r.NewRoute("PATCH", path, handlers)
}

func (r *RouteGroup) Head(path string, handlers ...HandlerFunc) {
	r.NewRoute("HEAD", path, handlers)
}

func (r *RouteGroup) Options(path string, handlers ...HandlerFunc) {
	r.NewRoute("OPTIONS", path, handlers)
}

func (r *RouteGroup) Connect(path string, handlers ...HandlerFunc) {
	r.NewRoute("CONNECT", path, handlers)
}

func (r *RouteGroup) Trace(path string, handlers ...HandlerFunc) {
	r.NewRoute("TRACE", path, handlers)
}