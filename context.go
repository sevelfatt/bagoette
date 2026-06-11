package bagoette

import (
	"errors"
	"log"

	"github.com/sevelfatt/bagoette/utils"
)

func NewContext() *Ctx {
	return &Ctx{
		currentHandlerIndex: 0,
		data: make(map[string]any),
	}
}

func (c *Ctx) Check() error {
	if c == nil {
		log.Println("Warning: Context is nil")
		return errors.New("Context is nil")
	}
	if c.currentRoute == nil {
		log.Println("Warning: Current route is nil")
		return errors.New("Current route is nil")
	}
	if c.w == nil {
		log.Println("Warning: Response writer is nil")
		return errors.New("Response writer is nil")
	}
	if c.r == nil {
		log.Println("Warning: Request is nil")
		return errors.New("Request is nil")
	}
	return nil
}

//Reset: reset the context
func (c *Ctx) Reset() error {
	err := c.Check()
	if err != nil {
		return err
	}
	c.currentHandlerIndex = 0
	c.data = nil
	return nil
}
//Set: set a value in the context data
func (c *Ctx) Set(key string, value any) error {
	err := c.Check()
	if err != nil {
		return err
	}
	if c.data == nil {
		c.data = make(map[string]any)
	}
	c.data[key] = value
	return nil
}

//Get: get a value from the context data
func (c *Ctx) Get(key string) (any, error) {
	err := c.Check()
	if err != nil {
		return nil, err
	}
	return c.data[key], nil
}

//Abort: abort the request
func (c *Ctx) Abort() error {
	err := c.Check()
	if err != nil {
		return err
	}
	c.currentHandlerIndex = len(c.currentRoute.Handlers)
	return nil
}

//Next: call the next handler and reset the context after the last handler is called
func (c *Ctx) Next() error {
	err := c.Check()
	if err != nil {
		return err
	}
	if c.currentHandlerIndex == len(c.currentRoute.Handlers) - 1 {
		//reset the context after the last handler is called
		c.Reset()
		return nil
	}
	//increment the handler index
	c.currentHandlerIndex++
	//call the next handler
	c.currentRoute.Handlers[c.currentHandlerIndex](c)
	return nil
}

//Bind: bind the request body to a struct
func (c *Ctx) Bind(body any) error {
	err := c.Check()
	if err != nil {
		return err
	}
	return utils.DecodeJSONFromRequestBody(c.r, body)
}

//Response: respond with a JSON
func (c *Ctx) Response(status int, data any) error {
	err := c.Check()
	if err != nil {
		return err
	}
	requestLog(c.r.Method, c.r.URL.Path, status)
	utils.RespondJSON(c.w, status, data)
	return nil
}

//Error: respond with an error
func (c *Ctx) Error(status int, message string) error {
	err := c.Check()
	if err != nil {
		return err
	}
	requestLog(c.r.Method, c.r.URL.Path, status)
	utils.RespondJSON(c.w, status, map[string]string{
		"error": message,
	})
	return nil
}

//Query: get a query parameter
func (c *Ctx) Query(key string) (string, error) {
	err := c.Check()
	if err != nil {
		return "", err
	}
	return c.r.URL.Query().Get(key), nil
}

//Header: get a header
func (c *Ctx) Header(key string) (string, error) {
	err := c.Check()
	if err != nil {
		return "", err
	}
	return c.r.Header.Get(key), nil
}

//SetHeader: set a header
func (c *Ctx) SetHeader(key string, value string) error {
	err := c.Check()
	if err != nil {
		return err
	}
	c.w.Header().Set(key, value)
	return nil
}

//Param: get a path parameter
func (c *Ctx) Param(key string) (string, error) {
	err := c.Check()
	if err != nil {
		return "", err
	}
	return c.r.PathValue(key), nil
}
