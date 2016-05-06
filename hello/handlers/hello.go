package handlers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Hello string
}

// Hello is a basic Handler for an Endpoint
func Hello(w http.ResponseWriter, r *http.Request) error {
	resp := response{Hello: "world"}

	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}
