package handlers

import (
	"encoding/json"
	"net/http"

	s "github.com/zknill/RESTume/service"
)

// Skills is a basic endpoint handler that exposes data about technical skills
func Skills(w http.ResponseWriter, r *http.Request) *s.HandlerError {
	// TODO: replace this with a document store instead of hardcoded data.
	var skills [2]string
	skills[0] = "Golang"
	skills[1] = "REST"
	b, err := json.Marshal(map[string]interface{}{
		"Skills": skills,
	})
	if err != nil {
		return s.NewError(err)
	}
	w.Write(b)
	return nil
}
