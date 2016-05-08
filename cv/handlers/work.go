package handlers

import (
	"net/http"
	"encoding/json"
	"time"
)

type Career struct {
	Jobs []*Job
}

type Job struct {
	Start int64
	End int64
	Company string
	Title string
	Notes string
}

func Work (w http.ResponseWriter, r *http.Request) error {
	// TODO: replace this with a document store instead of hardcoded data.

	c := Career{}
	notes := "Working on the launch of a web-based company . Working in an agile team, on the back-end of an" +
	"automatic document generation system, as well as writing requirements and designing the wireframes" +
	"for webpages to support the system."
	c.AddJob(&Job{
		Start: 1401580800,
		End: 1409529600,
		Company: "Geniac",
		Title: "IT Engineer/Developer",
		Notes: notes,
	})
	notes = "Using Agile methodologies to develop dockerized Go and Python micro-services with RESTful APIs." +
	"Including developing services to aggregate and process real-time data from servers all over the world."
	c.AddJob(&Job{
		Start: 1446336000,
		End: time.Now().Unix(),
		Company: "Multiplay",
		Title: "Developer",
		Notes: notes,
	})

	b, err := json.Marshal(c)
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}


func (c *Career) AddJob(j *Job) {
	c.Jobs = append(c.Jobs, j)
}