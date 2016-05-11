package main

import (
	"github.com/zknill/RESTume/cv/handlers"
	"github.com/zknill/RESTume/service"
	"github.com/zknill/RESTume/service/resources"
)

func main() {
	s := service.Init()

	s.AddEndpoint(&service.Endpoint{
		Name:        "About",
		Description: "A little about Zak",
		Route:       []string{"/about/"},
		Handle:      handlers.About,
		Methods:     []string{"GET"},
	})
	s.AddEndpoint(&service.Endpoint{
		Name:        "Hobbies",
		Description: "A little more about Zak and his interests",
		Route:       []string{"/hobbies/"},
		Handle:      handlers.Hobbies,
		Methods:     []string{"GET"},
	})
	s.AddEndpoint(&service.Endpoint{
		Name:        "Skills",
		Description: "Zak's technical skills",
		Route:       []string{"/skills/"},
		Handle:      handlers.Skills,
		Methods:     []string{"GET"},
	})
	s.AddEndpoint(&service.Endpoint{
		Name:        "Work",
		Description: "Zak's previous work experience",
		Route:       []string{"/work/", "/work/{company}"},
		Handle:      handlers.Work,
		Methods:     []string{"GET", "POST"},
	})

	career := &resources.Collection{
		Col:   "career",
		Index: []string{"Company"},
	}

	s.AddResource("db", &resources.Database{
		Name:        "tiedot",
		Location:    "/tmp/tiedot-database",
		Collections: []*resources.Collection{career},
	})

	s.Run()
}
