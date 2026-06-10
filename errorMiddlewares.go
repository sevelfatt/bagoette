package bagoette

import (
	"net/http"

	"github.com/sevelfatt/bagoette/utils"
)

func (R *Router) ApplyErrorMiddlewares(next HandlerFunc) HandlerFunc {
	return R.NotFoundMiddleware(R.MethodNotAllowedMiddleware(R.InternalServerErrorMiddleware(next)))
}

func (R *Router) NotFoundMiddleware(next HandlerFunc) HandlerFunc {
	fn := func(c *Context) {
		for _, route := range *R.routes {
			if utils.MatchRoute(route.path, c.r.URL.Path) {
				next(c)
				return
			}
		}
		c.Error(http.StatusNotFound, "Not Found")
	}
	return fn
}

func (R *Router) MethodNotAllowedMiddleware(next HandlerFunc) HandlerFunc {
	fn := func(c *Context) {
		for _, route := range *R.routes {
			if utils.MatchRoute(route.path, c.r.URL.Path) {
				next(c)
				return
			}
		}
		c.Error(http.StatusMethodNotAllowed, "Method Not Allowed")
	}
	return fn
}

func (R *Router) InternalServerErrorMiddleware(next HandlerFunc) HandlerFunc {
	fn := func(c *Context) {
		next(c)
		if c.r.Method != http.MethodGet {
			c.Error(http.StatusInternalServerError, "Internal Server Error")
		}
	}
	return fn
}