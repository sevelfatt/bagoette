package bagoette

import "net/http"

func (b *BagoetteClient) Get(path string, handler http.HandlerFunc) Route {
	pattern := "GET " + path
	return Route{
		pattern: pattern,
		handler: handler,
	}
	
}

func (b *BagoetteClient) Post(path string, handler http.HandlerFunc) Route {
	pattern := "POST " + path
	return Route{
		pattern: pattern,
		handler: handler,
	}
}

func (b *BagoetteClient) Put(path string, handler http.HandlerFunc) Route {
	pattern := "PUT " + path
	return Route{
		pattern: pattern,
		handler: handler,
	}
}

func (b *BagoetteClient) Delete(path string, handler http.HandlerFunc) Route {
	pattern := "DELETE " + path
	return Route{
		pattern: pattern,
		handler: handler,
	}
}

func (b *BagoetteClient) Patch(path string, handler http.HandlerFunc) Route {
	pattern := "PATCH " + path
	return Route{
		pattern: pattern,
		handler: handler,
	}
}

func (b *BagoetteClient) Head(path string, handler http.HandlerFunc) Route {
	pattern := "HEAD " + path
	return Route{
		pattern: pattern,
		handler: handler,
	}
}

func (b *BagoetteClient) Options(path string, handler http.HandlerFunc) Route {
	pattern := "OPTIONS " + path
	return Route{
		pattern: pattern,
		handler: handler,
	}
}

func (b *BagoetteClient) Trace(path string, handler http.HandlerFunc) Route {
	pattern := "TRACE " + path
	return Route{
		pattern: pattern,
		handler: handler,
	}
}

func (b *BagoetteClient) Connect(path string, handler http.HandlerFunc) Route {
	pattern := "CONNECT " + path
	return Route{
		pattern: pattern,
		handler: handler,
	}
}