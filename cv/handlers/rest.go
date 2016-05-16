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
	colName := vars["collection"]

	idx := vars["index"]
	val := vars["value"]

	eh := context.Get(r, s.ContextKey).(*s.EndpointHandler)
	d := eh.Resources["db"].(*db.Database)
	cols := d.Collections

	c, err := validateCol(cols, colName)
	if err != nil {
		return err
	}

	if !validateIdx(c, idx) {
		return &s.HandlerError{Err: fmt.Errorf("Not found: index does not exist."), Code: 404}
	}

	q := map[string]interface{}{
		"eq": val,
		"in": []string{idx},
	}

	dbCol := d.Data.Use(c.Name)

	b, dbErr := db.Query(dbCol, q)
	if dbErr != nil {
		return &s.HandlerError{Err: fmt.Errorf("Not found."), Code: 404}
	}

	w.Write(b)
	return nil
}

func validateCol(cols []*db.Collection, col string) (*db.Collection, *s.HandlerError) {
	for _, c := range cols {
		if c.Name == col {
			return c, nil
		}
	}
	return nil, &s.HandlerError{Err: fmt.Errorf("Not found: collection does not exist."), Code: 404}
}

func validateIdx(col *db.Collection, idx string) bool {
	for _, i := range col.Index {
		if i == idx {
			return true
		}
	}
	return false
}
