package handlers

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/golang/glog"
	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zknill/RESTume/service"
	db "github.com/zknill/RESTume/service/database"
)

func TestWork(t *testing.T) {
	loc, handler := setup()
	defer os.RemoveAll(loc)

	ts := httptest.NewServer(handler)
	defer ts.Close()

	type job map[string]interface{}
	data := job{
		"Company": "Geniac",
		"Notes":   "Working on the launch of a web-based company",
		"Title":   "IT Engineer/Developer",
	}

	b, _ := json.Marshal(data)

	body := bytes.NewReader(b)

	http.Post(ts.URL+"/work/", "application/json", body)

	resp, err := http.Get(ts.URL + "/work/Geniac")
	if err != nil {
		defer resp.Body.Close()
	}

	rdata, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		glog.Error(err)
	}

	check := make(map[string]job)
	json.Unmarshal(rdata, &check)

	for _, j := range check {
		assert.Equal(t, j, data)
	}

}

// TODO: build a proper test rig so that this setup isn't needed per end-point that we test.
func setup() (string, http.Handler) {
	s := service.Init()

	s.AddEndpoint(&service.Endpoint{
		Name:        "Work",
		Description: "Zak's previous work experience",
		Route:       []string{"/work/", "/work/{company}"},
		Handle:      Work,
		Methods:     []string{"GET", "POST"},
	})

	career := &db.Collection{
		Name:  "career",
		Index: []string{"Company"},
	}

	loc := "/tmp/tiedot-test-database"
	s.AddResource("db", &db.Database{
		Name:        "tiedot",
		Location:    loc,
		Collections: []*db.Collection{career},
	})

	for _, r := range s.Resources {
		r.Init()
	}

	router := mux.NewRouter()
	// Register all the endpoints
	for _, e := range s.Endpoints {
		// Register the different routes for each endpoint
		for _, r := range e.Route {
			router.Handle(r, service.NewEndpointHandler(e, s.Resources)).Methods(e.Methods...)
		}
	}

	return loc, router
}
