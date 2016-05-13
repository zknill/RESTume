package handlers

import (
	"fmt"
	"net/http"

	tiedot "github.com/HouzuoGuo/tiedot/db"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	s "github.com/zknill/RESTume/service"
	db "github.com/zknill/RESTume/service/database"
)

//TODO: refactor this file, try and abstract database references
//TODO: add authentication to this endpoint. Anyone can add data into the db.

// Work is the work endpoint handler
func Work(w http.ResponseWriter, req *http.Request) error {
	// TODO: architect the context and resources better - this is horrible.
	sh := context.Get(req, s.ContextKey).(*s.Handler)
	data := sh.Resources["db"].(*db.Database).Data
	col := data.Use("career")

	// If we are POSTing data
	if req.Method == "POST" {
		if _, err := db.Insert(req, col); err != nil {
			return fmt.Errorf("Database insert error: %s", err)
		}
		w.Write([]byte(`{"success": "True"}`))
		return nil
	}

	var resp []byte
	var err error
	vars := mux.Vars(req)
	// If we are getting a specific Company
	if comp, ok := vars["company"]; ok {
		q := map[string]interface{}{
			"eq": comp,
			"in": []string{"Company"},
		}
		resp, err = db.Query(col, q)
	} else {
		// Else we are getting all the docs
		queryResult := &map[int]struct{}{}
		tiedot.EvalAllIDs(col, queryResult)
		resp, err = db.FlatResult(col, queryResult)
	}

	w.Write(resp)
	return err
}
