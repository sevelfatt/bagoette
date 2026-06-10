package bagoette

import "net/http"

func (R *Router) NotFoundMiddleware(next http.HandlerFunc) http.HandlerFunc {
	fn := func(w http.ResponseWriter, r *http.Request) {
		for _, route := range *R.routes {
			if route.path == r.URL.Path {
				next(w, r)
				return
			}
		}
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("404 Not Found"))
	}
	return fn
}