# [Development 🛠️] **BAGOETTE**: Best Go RestfulAPI Library for Beginner 🥖
![Teto want to say hello](https://images.steamusercontent.com/ugc/2290709938774550184/4756C519C526F8CD86AD897B080007537894E64A/?imw=268&imh=268&ima=fit&impolicy=Letterbox&imcolor=%23000000&letterbox=true)
## yo whats up gng ✌️
So currently i'am trying to construct a **Golang RestfulAPI** library like **gin**, **fiber** and **echo** 🏗️ but... its specialized to help beginner to create their own Golang backend service.

yeah, i know the codes and the project structures are still really messy because i'm new to this Programming language🥀

so yeah maybe you guys can give me some advice?

## documentation
### installation
just give it a try dawg 🥖
```bash
go get github.com/sevelfatt/bagoette
```

### quick example
```go
package main

import "github.com/sevelfatt/bagoette"

func main() {
	b := bagoette.NewClient() // create new client
	r := b.NewRouter() // create new router

	r.Get("/ping", func(c *bagoette.Ctx){
		c.Respond(200, "pong")
	}) // open new route with GET method on /ping path

	b.Serve() // start the server
}
```

### use middleware
```go
package main

import (
	"github.com/sevelfatt/bagoette"
)

func main() {
	b := bagoette.NewClient()
	r := b.NewRouter()

	r.Get("/ping", loggerMiddleware, pong) // you can add more middleware like this

	b.Serve()
}

func loggerMiddleware(c *bagoette.Ctx) {
	bagoette.Logger.Println("Request: " + c.request.Method + " " + c.request.URL.Path)
	c.Next()
}

func pong(c *bagoette.Ctx) {
	c.Respond(200, "pong")
}
```

### use router group
```go
package main

import "github.com/sevelfatt/bagoette"

func main() {
	b := bagoette.NewClient()
	r := b.NewRouter()

	userRouter := r.Group("/users") //you can wrap some routes with same prefix like this
	userRouter.Get("/ping", pong) //so the path will be /users/ping

	b.Serve()
}

func pong(c *bagoette.Ctx) {
	c.Respond(200, "pong")
}
```

### use middleware in router group
```go
package main

import (
	"github.com/sevelfatt/bagoette"
)

func main() {
	b := bagoette.NewClient()
	r := b.NewRouter()

	userRouter := r.Group("/users") 
    userRouter.Use(loggerMiddleware) // you can add more middleware like this

	userRouter.Get("/ping", pong)

	b.Serve()
}

func loggerMiddleware(c *bagoette.Ctx) {
	bagoette.Logger.Println("Request: " + c.request.Method + " " + c.request.URL.Path)
	c.Next()
}

func pong(c *bagoette.Ctx) {
	c.Respond(200, "pong")
}
```


