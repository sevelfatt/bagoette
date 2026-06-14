package bagoette

import (
	"errors"

	"github.com/sevelfatt/bagoette/utils"
)

//Router struct: work as the router of the server
type RouteGroup struct {
	routes *[]Route
	middlewares []HandlerFunc
	prefix string
}

//Route struct: Define individual route
type Route struct {
	path    string
	method  string

	pathSegments []string
	paramKeys []string

	handlers []HandlerFunc
}

func (r *Route) Check() error {
	if r == nil {
		return errors.New("Can't add nil route")
	}
	if r.handlers == nil {
		return errors.New("Can't add route with no handlers")
	}
	if r.method == "" {
		return errors.New("Can't add route with no method")
	}
	if r.path == "" {
		return errors.New("Can't add route with no path")
	}
	if r.pathSegments == nil {
		return errors.New("Can't add route with no path segments")
	}
	return nil
}

func (b *BagoetteClient) NewRouter() *RouteGroup {
	return &RouteGroup{
		routes:      b.routes,
		middlewares: []HandlerFunc{b.NotFoundMiddleware, b.MethodNotAllowedMiddleware, b.InternalServerErrorMiddleware},
		prefix:      "",
	}
}

func (r *RouteGroup) Group(prefix string) *RouteGroup {
	newRouter := *r
	newRouter.prefix = r.prefix + prefix
	return &newRouter
}

func (r *RouteGroup) Use(handlers ...HandlerFunc) *RouteGroup {
	r.middlewares = append(r.middlewares, handlers...)
	return r
}

func (r *RouteGroup) NewRoute(method string, path string, handlers []HandlerFunc) error {
	fullPath := r.prefix + path
	
	pathSegments := utils.GetPathSegment(fullPath)
	paramKeys := utils.GetParamKeys(pathSegments)

	route := &Route{
		method:   method,
		path:     fullPath,
		pathSegments: pathSegments,
		paramKeys: paramKeys,
		handlers: append(r.middlewares, handlers...),
	}
	return r.AddNewRouteToBagoetteClient(route)
}

func (r *RouteGroup) AddNewRouteToBagoetteClient(route *Route) error {
	err := route.Check()
	if err != nil {
		logger.Println("Error adding route:", err)
		return err
	}
	*r.routes = append(*r.routes, *route)
	return nil
}