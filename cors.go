package bagoette

type Cors struct {
	AllowedOrigins []string
	AllowedMethods []string
	AllowedHeaders []string
	ExposedHeaders []string
	MaxAge int
	AllowCredentials bool
}

func NewCors() *Cors {
	return &Cors{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{"*"},
		AllowedHeaders: []string{"*"},
		ExposedHeaders: []string{"*"},
		MaxAge: 0,
		AllowCredentials: false,
	}
}

func (c *Cors) SetAllowedOrigins(origins []string) *Cors {
	c.AllowedOrigins = origins
	return c
}

func (c *Cors) SetAllowedMethods(methods []string) *Cors {
	c.AllowedMethods = methods
	return c
}

func (c *Cors) SetAllowedHeaders(headers []string) *Cors {
	c.AllowedHeaders = headers
	return c
}

func (c *Cors) SetExposedHeaders(headers []string) *Cors {
	c.ExposedHeaders = headers
	return c
}

func (c *Cors) SetMaxAge(maxAge int) *Cors {
	c.MaxAge = maxAge
	return c
}

func (c *Cors) SetAllowCredentials(allowCredentials bool) *Cors {
	c.AllowCredentials = allowCredentials
	return c
}