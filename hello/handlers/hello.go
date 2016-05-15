package handlers

import (
	"encoding/json"
	"net/http"

	s "github.com/zknill/RESTume/service"
)

type response struct {
	Hello string
}

// Hello is a basic Handler for an Endpoint
func Hello(w http.ResponseWriter, r *http.Request) *s.HandlerError {
	resp := response{Hello: "world"}

	b, err := json.Marshal(resp)
	if err != nil {
		return s.NewError(err)
	}
	w.Write(b)
	return nil
}
