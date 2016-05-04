package main

import (
	"log"
	"net/http"
	"github.com/zknill/RESTume/hello/handlers"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/", handlers.Hello)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
