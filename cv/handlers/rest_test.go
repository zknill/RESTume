package handlers

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"os"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/zknill/RESTume/service"
	db "github.com/zknill/RESTume/service/database"
)

func TestREST(t *testing.T) {
	loc, r := setupREST()
	ts := httptest.NewServer(r)
	defer func(ts *httptest.Server, a string) {
		ts.Close()
		os.RemoveAll(loc)
	}(ts, loc)

	resp, _ := http.Get(ts.URL + "/abc/")

	b, _ := ioutil.ReadAll(resp.Body)

	assert.Equal(t, []byte("Not found\n"), b)
}

func setupREST() (string, http.Handler) {
	s := service.Init()

	s.AddEndpoint(&service.Endpoint{
		Name:        "REST",
		Description: "RESTful endpoint that only uses the URL",
		Route:       []string{"/{collection}/", "/{collection}/{index}/{field}/{value}"},
		Handle:      REST,
		Methods:     []string{"GET", "POST"},
	})

	career := &db.Collection{
		Col:   "career",
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
