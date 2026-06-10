package bagoette

import "net/http"

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func (c *Context) Write(b []byte) (int, error) {
	return c.w.Write(b)
}

func (c *Context) Read(b []byte) (int, error) {
	return c.r.Body.Read(b)
}