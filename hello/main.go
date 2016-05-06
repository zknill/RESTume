package main

import (
	"github.com/zknill/RESTume/service"
	"github.com/zknill/RESTume/hello/handlers"
)

func main() {
	s := service.Init()
	end := service.NewEndpoint(
		"Hello World",
		"A very endpoint to test the service implementation",
		"/",
		handlers.Hello,
	)

	s.AddEndpoint(end)
	s.Run()
}
