package bagoette

func FirstHandlerMiddleware(c *Context) {
	c.Reset()
	c.Next()
}