package bagoette

import (
	"net/http"
	"runtime/debug"
	"strconv"
	"strings"

	"github.com/sevelfatt/bagoette/utils"
)

func (b *BagoetteClient) CorsMiddleware(c *Ctx) {
		c.request.Header.Set("Access-Control-Allow-Origin", strings.Join(b.Opts.Cors.AllowedOrigins, ", "))
		c.request.Header.Set("Access-Control-Allow-Methods", strings.Join(b.Opts.Cors.AllowedMethods, ", "))
		c.request.Header.Set("Access-Control-Allow-Headers", strings.Join(b.Opts.Cors.AllowedHeaders, ", "))
		c.request.Header.Set("Access-Control-Expose-Headers", strings.Join(b.Opts.Cors.ExposedHeaders, ", "))
		c.request.Header.Set("Access-Control-Max-Age", strconv.Itoa(b.Opts.Cors.MaxAge))
		c.request.Header.Set("Access-Control-Allow-Credentials", strconv.FormatBool(b.Opts.Cors.AllowCredentials))
		if c.request.Method == http.MethodOptions {
			c.writer.WriteHeader(http.StatusOK)
			return
		}
		c.Next()
}

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