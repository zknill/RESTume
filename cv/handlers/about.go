package handlers

import (
	"net/http"
	"encoding/json"
)

func About (w http.ResponseWriter, r *http.Request) error {
	// TODO: replace this with a document store instead of hardcoded data.
	b, err := json.Marshal(map[string]interface{}{
		"Name": "Zak",
		"Job": "Developer",
	})
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}
