package handlers

import (
	"encoding/json"
	"net/http"

	s "github.com/zknill/RESTume/service"
)

// Hobbies is a basic endpoint handler that exposes data about interests.
func Hobbies(w http.ResponseWriter, r *http.Request) *s.HandlerError {
	// TODO: replace this with a document store instead of hardcoded data.
	hobbies := [3]string{"Running", "Marathons", "Code"}
	b, err := json.Marshal(map[string]interface{}{
		"Hobbies": hobbies,
	})
	if err != nil {
		return s.NewError(err)
	}
	w.Write(b)
	return nil
}
