package handlers

import (
	"fmt"
	"net/http"

	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	s "github.com/zknill/RESTume/service"
	db "github.com/zknill/RESTume/service/database"
)

// REST is the generic endpoint handler that relies on RESTful URLs
func REST(w http.ResponseWriter, r *http.Request) *s.HandlerError {
	// RESTful URL/{collection}/{index}/{query}

	vars := mux.Vars(r)
	col := vars["collection"]

	//idx := vars["index"]
	//field := vars["field"]
	//val := vars["val"]

	eh := context.Get(r, s.ContextKey).(*s.EndpointHandler)
	cols := eh.Resources["db"].(*db.Database).Collections

	bCol := false
	for _, c := range cols {
		if c.Col == col {
			bCol = true
		}
	}
	if !bCol {
		return &s.HandlerError{
			Err:  fmt.Errorf("Not found"),
			Code: http.StatusNotFound,
		}
	}
	return nil
}
