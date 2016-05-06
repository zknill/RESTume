package service

import (
	"net/http"

	log "github.com/golang/glog"
	"github.com/gorilla/mux"
)

// Service represents a web-service
type Service struct {
	Name      string
	Endpoints []*Endpoint
}

// Endpoint represents a single API endpoint
type Endpoint struct {
	Name        string
	Description string
	Route       string
	Handle      Handler
}

// ServiceHandler allows implements the http.Handler interface
type ServiceHandler struct {
	Endpoint *Endpoint
}

// Handler is a func that will handle endpoint requests
type Handler func(http.ResponseWriter, *http.Request) error

// Init returns a new empty web-service
func Init() *Service {
	return &Service{}
}

// Run kicks off the service and adds all the handlers for the endpoints.
func (s *Service) Run() {
	router := mux.NewRouter()
	for _, e := range s.Endpoints {
		router.Handle(e.Route, NewServiceHandler(e))
	}
	log.Fatal(http.ListenAndServe(":8000", nil))
}

// AddEndpoint adds the given endpoint to the service
func (s *Service) AddEndpoint(e *Endpoint) {
	s.Endpoints = append(s.Endpoints, e)
}

// NewEndpoint takes the details to create a new endpoint and returns a new endpoint
func NewEndpoint(name, desc, route string, handle Handler) *Endpoint {
	return &Endpoint{
		Name:        name,
		Description: desc,
		Route:       route,
		Handle:      handle,
	}
}

// NewServiceHandler returns a new Handler for the service that implements http.Handler
func NewServiceHandler(end *Endpoint) *ServiceHandler {
	return &ServiceHandler{Endpoint: end}
}

// ServeHTTP implements the http.Handler interface for a ServiceHandler
func (sh *ServiceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if err := sh.Endpoint.Handle(w, r); err != nil {
		log.Error(err.Error())
		http.Error(w, err.Error(), 500)
	}
}
