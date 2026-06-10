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