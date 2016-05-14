package service

import (
	"net/http"

	log "github.com/golang/glog"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

// Service represents a web-service
type Service struct {
	Name      string
	Endpoints []*Endpoint
	Resources map[string]Resource
}

// Endpoint represents a single API endpoint
type Endpoint struct {
	Name        string
	Description string
	Route       []string
	Handle      Handle
	Methods     []string
}

// Handler implements the http.Handler interface
type Handler struct {
	Endpoint  *Endpoint
	Resources map[string]Resource
}

// Handle is a func that will handle endpoint requests, e.g. Work
type Handle func(http.ResponseWriter, *http.Request) error

type key int

// ContextKey is the key used to look up details in teh gorilla context
const ContextKey key = 0

// Init returns a new empty web-service
func Init() *Service {
	s := &Service{}
	s.Resources = make(map[string]Resource)
	return s
}

// Run kicks off the service and adds all the handlers for the endpoints.
func (s *Service) Run() {
	log.Infoln("Running " + s.Name)
	// Setup all the resources that the service requires.
	for _, r := range s.Resources {
		r.Init()
	}

	router := mux.NewRouter()
	// Register all the endpoints
	for _, e := range s.Endpoints {
		// Register the different routes for each endpoint
		for _, r := range e.Route {
			router.Handle(r, NewServiceHandler(e, s.Resources)).Methods(e.Methods...)
		}
	}

	// Register the router as the http.Handler
	http.Handle("/", Logger(router, s.Name))
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// AddEndpoint adds the given endpoint to the service
func (s *Service) AddEndpoint(e *Endpoint) {
	s.Endpoints = append(s.Endpoints, e)
}

// AddResource adds the required external resources to the service
func (s *Service) AddResource(key string, r Resource) {
	s.Resources[key] = r
}

// NewServiceHandler returns a new Handler for the service that implements http.Handler
func NewServiceHandler(end *Endpoint, r map[string]Resource) *Handler {
	return &Handler{Endpoint: end, Resources: r}
}

// ServeHTTP implements the http.Handler interface for a ServiceHandler
func (sh *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// all responses are JSON.
	w.Header().Add("Content-type", "application/json")

	// Set up a context so the endpoints can access the resources
	context.Set(r, ContextKey, sh)

	// Call the correct handler for the endpoint
	if err := sh.Endpoint.Handle(w, r); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), 500)
	}
}
