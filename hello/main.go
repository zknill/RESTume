package main

import (
	"github.com/zknill/RESTume/hello/handlers"
	"github.com/zknill/RESTume/service"
)

func main() {
	s := service.Init()
	end := &service.Endpoint{
		Name:        "Hello World",
		Description: "A very simple endpoint to test the service implementation",
		Route:       "/",
		Handle:      handlers.Hello,
	}

	s.AddEndpoint(end)
	s.Run()
}
