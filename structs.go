package bagoette

import (
	"net/http"
)

//Bagoette main struct: work as the core of the library
//and provide all the features like router, middleware, context, etc
type BagoetteClient struct {
	opts *BagoetteOptions

	httpClient *http.Client
	httpHandler *http.ServeMux

	routes *[]Route
}

type BagoetteOptions struct {
	Port int
	Host string
	
}

//Context struct: work as the container of the request and response
type Ctx struct {
	writer http.ResponseWriter
	request *http.Request

	currentHandlerIndex int
	currentRoute *Route

	data map[string]any
}

//HandlerFunc type: Define the handler function that use bagoette context
type HandlerFunc func(c *Ctx)

//Router struct: work as the router of the server
type Router struct {
	httpHandler *http.ServeMux
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