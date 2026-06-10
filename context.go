package bagoette

import (
	"net/http"

	"github.com/sevelfatt/bagoette/utils"
)

type Context struct {
	w http.ResponseWriter
	r *http.Request
}

func (c *Context) Bind(body any) error {
	return utils.DecodeJSON(c.r, body)
}

func (c *Context) Response(status int, data any) {
	utils.RespondJSON(c.w, status, data)
}

func (c *Context) Error(status int, message string) {
	utils.RespondJSON(c.w, status, map[string]string{
		"error": message,
	})
}

func (c *Context) Query(key string) string {
	return c.r.URL.Query().Get(key)
}

func (c *Context) Header(key string) string {
	return c.r.Header.Get(key)
}

func (c *Context) SetHeader(key string, value string) {
	c.w.Header().Set(key, value)
}

func (c *Context) Param(key string) string {
	return c.r.PathValue(key)
}
