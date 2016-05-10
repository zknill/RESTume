package handlers

import (
	"encoding/json"
	"net/http"
)

func Skills(w http.ResponseWriter, r *http.Request) error {
	// TODO: replace this with a document store instead of hardcoded data.
	var skills [2]string
	skills[0] = "Golang"
	skills[1] = "REST"
	b, err := json.Marshal(map[string]interface{}{
		"Skills": skills,
	})
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}
