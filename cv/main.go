package main

import (
	"flag"

	"github.com/zknill/RESTume/cv/handlers"
	"github.com/zknill/RESTume/service"
	db "github.com/zknill/RESTume/service/database"
)

func main() {
	flag.Parse()
	s := service.Init()
	s.Name = "cv"

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
	s.AddEndpoint(&service.Endpoint{
		Name:        "REST",
		Description: "RESTful endpoint that only uses the URL",
		Route:       []string{"/{collection}/", "/{collection}/{index}/{field}/{value}"},
		Handle:      handlers.REST,
		Methods:     []string{"GET"},
	})

	career := &db.Collection{
		Col:   "career",
		Index: []string{"Company"},
	}

	s.AddResource("db", &db.Database{
		Name:        "tiedot",
		Location:    "/tmp/tiedot-database",
		Collections: []*db.Collection{career},
	})

	s.Run()
}
