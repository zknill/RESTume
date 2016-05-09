package main

import (
	"github.com/zknill/RESTume/cv/handlers"
	"github.com/zknill/RESTume/service"
)

func main() {
	s := service.Init()
	s.AddEndpoint(&service.Endpoint{
		Name:        "About",
		Description: "A little about Zak",
		Route:       "/about/",
		Handle:      handlers.About,
	})
	s.AddEndpoint(&service.Endpoint{
		Name:        "Hobbies",
		Description: "A little more about Zak and his interests",
		Route:       "/hobbies/",
		Handle:      handlers.Hobbies,
	})
	s.AddEndpoint(&service.Endpoint{
		Name:        "Skills",
		Description: "Zak's technical skills",
		Route:       "/skills/",
		Handle:      handlers.Skills,
	})
	s.AddEndpoint(&service.Endpoint{
		Name:        "Work",
		Description: "Zak's previous work experience",
		Route:       "/work/",
		Handle:      handlers.Work,
	})

	s.AddResource("db", &service.Database{
		Name:     "tiedot",
		Location: "/tmp/database",
	})

	s.Run()
}
