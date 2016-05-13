package handlers

import (
	"encoding/json"
	"net/http"
)

// Hobbies is a basic endpoint handler that exposes data about interests.
func Hobbies(w http.ResponseWriter, r *http.Request) error {
	// TODO: replace this with a document store instead of hardcoded data.
	hobbies := [3]string{"Running", "Marathons", "Code"}
	b, err := json.Marshal(map[string]interface{}{
		"Hobbies": hobbies,
	})
	if err != nil {
		return err
	}
	w.Write(b)
	return nil
}
