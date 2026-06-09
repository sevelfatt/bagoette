package bagoette

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Route struct {
	Name string
	Path string
	Handler http.HandlerFunc
	MuxRoute *mux.Route
}


func (r *Route) SetName(name string) *Route {
	r.Name = name
	return r
}

func (b *BagoetteClient) AddRoute(route Route) {
	route.MuxRoute = b.MuxRouter.Handle(route.Path, route.Handler)
	b.Routes = append(b.Routes, &route)
}
