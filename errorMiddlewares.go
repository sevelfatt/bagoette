package bagoette

import (
	"net/http"

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
			if utils.MatchRouteMethod(route.method, c.request) {
				c.Next()
				return
			}
		}
		c.Error(http.StatusMethodNotAllowed, "Method Not Allowed")
}

func (b *BagoetteClient) InternalServerErrorMiddleware(c *Ctx) {
	c.Next()
	if c.request.Method != http.MethodGet {
		c.Error(http.StatusInternalServerError, "Internal Server Error")
		return
	}
}