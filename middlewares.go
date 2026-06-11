package bagoette

func FirstHandlerMiddleware(c *Ctx) {
	c.Reset()
	c.Next()
}