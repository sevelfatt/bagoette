package bagoette

import "net/http"

func (R *Router) NotFoundMiddleware(next HandlerFunc) HandlerFunc {
	fn := func(c *Context) {
		for _, route := range *R.routes {
			if route.path == c.r.URL.Path {
				next(c)
				return
			}
		}
		c.w.WriteHeader(http.StatusNotFound)
		c.w.Write([]byte("404 Not Found"))
	}
	return fn
}