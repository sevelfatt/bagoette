package bagoette

import (
	"net/http"
)

func (b *BagoetteClient) registerAllRoutes() error {
	for _, route := range *b.routes {
		err := b.registerRoute(&route)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *BagoetteClient) registerRoute(route *Route) error {
	err := route.Check()
	if err != nil {
		logger.Println("Error registering route:", err)
		return err
	}
	fn := func(w http.ResponseWriter, r *http.Request){
		context := NewContext(b.Opts, w, r, route)
		context.currentRoute.handlers[0](context)
	}

	b.httpHandler.Handle(route.method + " " + route.path, http.HandlerFunc(fn))
	return nil
}
