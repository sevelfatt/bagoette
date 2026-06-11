package bagoette

import "net/http"

//Bagoette main struct: work as the core of the library
//and provide all the features like router, middleware, context, etc
type BagoetteClient struct {
	port int

	httpClient *http.Client
	httpHandler *http.ServeMux
	
	routes *[]Route
	router *Router
	context *Context
}

type Router struct {
	httpHandler *http.ServeMux
	context *Context
	routes *[]Route
	middlewares []HandlerFunc
	prefix string
}
type HandlerFunc func(c *Context)

type Route struct {
	name    string
	path    string
	method  string
	pattern string
	Handlers []HandlerFunc
}

//Context struct: work as the container of the request and response
type Context struct {
	w http.ResponseWriter
	r *http.Request

	currentHandlerIndex int
	currentRoute *Route

	data map[string]any
}
