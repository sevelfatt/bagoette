package bagoette

import "net/http"

func (b *BagoetteClient) Get(path string, handler http.HandlerFunc) Route {
	fullPath := "GET " + path
	return Route{
		Path: fullPath,
		Handler: handler,
	}
	
}

func (b *BagoetteClient) Post(path string, handler http.HandlerFunc) Route {
	fullPath := "POST " + path
	return Route{
		Path: fullPath,
		Handler: handler,
	}
}

func (b *BagoetteClient) Put(path string, handler http.HandlerFunc) Route {
	fullPath := "PUT " + path
	return Route{
		Path: fullPath,
		Handler: handler,
	}
}

func (b *BagoetteClient) Delete(path string, handler http.HandlerFunc) Route {
	fullPath := "DELETE " + path
	return Route{
		Path: fullPath,
		Handler: handler,
	}
}

func (b *BagoetteClient) Patch(path string, handler http.HandlerFunc) Route {
	fullPath := "PATCH " + path
	return Route{
		Path: fullPath,
		Handler: handler,
	}
}

func (b *BagoetteClient) Head(path string, handler http.HandlerFunc) Route {
	fullPath := "HEAD " + path
	return Route{
		Path: fullPath,
		Handler: handler,
	}
}

func (b *BagoetteClient) Options(path string, handler http.HandlerFunc) Route {
	fullPath := "OPTIONS " + path
	return Route{
		Path: fullPath,
		Handler: handler,
	}
}

func (b *BagoetteClient) Trace(path string, handler http.HandlerFunc) Route {
	fullPath := "TRACE " + path
	return Route{
		Path: fullPath,
		Handler: handler,
	}
}

func (b *BagoetteClient) Connect(path string, handler http.HandlerFunc) Route {
	fullPath := "CONNECT " + path
	return Route{
		Path: fullPath,
		Handler: handler,
	}
}