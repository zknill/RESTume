package handlers

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Hello string
}

func Hello(w http.ResponseWriter, r *http.Request) {
	resp := response{Hello: "world"}

	// TODO(zak): don't swallow the error
	b, _ := json.Marshal(resp)
	w.Write(b)
}
