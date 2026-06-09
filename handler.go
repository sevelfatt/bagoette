package bagoette

import "net/http"

func (b *BagoetteClient) Get(path string, handler http.HandlerFunc) {
	b.MuxRouter.Handle("GET " + path, handler)
}

func (b *BagoetteClient) Post(path string, handler http.HandlerFunc) {
	b.MuxRouter.Handle("POST " + path, handler)
}

func (b *BagoetteClient) Put(path string, handler http.HandlerFunc) {
	b.MuxRouter.Handle("PUT " + path, handler)
}

func (b *BagoetteClient) Delete(path string, handler http.HandlerFunc) {
	b.MuxRouter.Handle("DELETE " + path, handler)
}

func (b *BagoetteClient) Patch(path string, handler http.HandlerFunc) {
	b.MuxRouter.Handle("PATCH " + path, handler)
}

func (b *BagoetteClient) Head(path string, handler http.HandlerFunc) {
	b.MuxRouter.Handle("HEAD " + path, handler)
}

func (b *BagoetteClient) Options(path string, handler http.HandlerFunc) {
	b.MuxRouter.Handle("OPTIONS " + path, handler)
}

func (b *BagoetteClient) Trace(path string, handler http.HandlerFunc) {
	b.MuxRouter.Handle("TRACE " + path, handler)
}

func (b *BagoetteClient) Connect(path string, handler http.HandlerFunc) {
	b.MuxRouter.Handle("CONNECT " + path, handler)
}