package main

import (
	"net/http"
	"github.com/zknill/RESTume/hello/handlers"
)

func main() {
	http.HandleFunc("/", handlers.Hello)
	http.ListenAndServe(":8000", nil)
}
