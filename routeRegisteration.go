package bagoette

func (b *BagoetteClient) registerAllRoutes() {
	for _, route := range *b.routes {
		b.RegisterRoute(&route)
	}
}

func (b *BagoetteClient) RegisterRoute(route *Route) {
	b.httpHandler.Handle(route.pattern, route.handler)
}
