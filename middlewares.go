package bagoette

import (
	"net/http"
	"runtime/debug"

	"github.com/sevelfatt/bagoette/utils"
)

func (b *BagoetteClient) NotFoundMiddleware(c *Ctx) {
	for _, route := range *b.routes {
		if utils.MatchRoute(route.pathSegments, utils.GetPathSegment(c.request.URL.Path)) {
			c.Next()
			return
		}
	}
	c.Error(http.StatusNotFound, "Not Found")
}

func (b *BagoetteClient) MethodNotAllowedMiddleware(c *Ctx) {
	for _, route := range *b.routes {
		if utils.MatchRoute(route.pathSegments, utils.GetPathSegment(c.request.URL.Path)) && utils.MatchRouteMethod(route.method, c.request) {
			c.Next()
			return
		}
	}
	c.Error(http.StatusMethodNotAllowed, "Method Not Allowed")
}

func (b *BagoetteClient) InternalServerErrorMiddleware(c *Ctx) {
	defer func() {
		if err := recover(); err != nil {
			// Log the panic with a stack trace
			Logger.Println("panic recovered",
				"error", err,
				"stack", string(debug.Stack()),
				"path", c.request.URL.Path,
			)
			c.Error(http.StatusInternalServerError, "internal server error")
		}
	}()
	c.Next()

}