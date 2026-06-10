package bagoette

import (
	"net/http"

	"github.com/sevelfatt/bagoette/utils"
)

//Context struct: work as the container of the request and response
type Context struct {
	w http.ResponseWriter
	r *http.Request

	currentHandlerIndex int
	handlers []HandlerFunc
	data map[string]any
}

func (c *Context) Reset() {
	c.currentHandlerIndex = 0
	c.handlers = nil
	c.data = nil
}

func (c *Context) Set(key string, value any) {
	if c.data == nil {
		c.data = make(map[string]any)
	}
	c.data[key] = value
}

func (c *Context) Get(key string) any {
	return c.data[key]
}

func (c *Context) Abort() {
	c.currentHandlerIndex = len(c.handlers)
}

func (c *Context) Next() {
	if c.currentHandlerIndex == len(c.handlers) {
		return
	}
	c.currentHandlerIndex++
	c.handlers[c.currentHandlerIndex](c)
}

func (c *Context) Bind(body any) error {
	return utils.DecodeJSONFromRequestBody(c.r, body)
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
