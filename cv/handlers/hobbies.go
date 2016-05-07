package handlers

import (
	"net/http"
	"encoding/json"
)

func Hobbies (w http.ResponseWriter, r *http.Request) error {
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

