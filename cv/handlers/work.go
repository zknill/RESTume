package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	tiedot "github.com/HouzuoGuo/tiedot/db"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	s "github.com/zknill/RESTume/service"
	resc "github.com/zknill/RESTume/service/resources"
)

//TODO: refactor this file, try and abstract database references
//TODO: add authentication to this endpoint. Anyone can add data into the db.

// Work is the work endpoint handler
func Work(w http.ResponseWriter, r *http.Request) error {
	// TODO: architect the context and resources better - this is horrible.
	sh := context.Get(r, s.ContextKey).(*s.ServiceHandler)
	db := sh.Resources["db"].(*resc.Database).Data
	col := db.Use("career")

	// If we are POSTing data
	if r.Method == "POST" {
		if _, err := resc.DBInsert(r, col); err != nil {
			return fmt.Errorf(fmt.Sprint("Database insert error: %s", err))
		}
		w.Write([]byte(`{"success": "True"}`))
		return nil
	}

	var resp []byte
	var err error
	vars := mux.Vars(r)
	// If we are getting a specific Company
	if comp, ok := vars["company"]; ok {
		q := map[string]interface{}{
			"eq": comp,
			"in": []string{"Company"},
		}
		resp, err = doQuery(col, q)
	} else {
		// Else we are getting all the docs
		queryResult := &map[int]struct{}{}
		tiedot.EvalAllIDs(col, queryResult)
		resp, err = flatRes(col, queryResult)
	}

	w.Write(resp)
	return err
}

func doQuery(col *tiedot.Col, q map[string]interface{}) (resp []byte, err error) {
	var query interface{}

	// TODO: Find a better way to do this. Marshal and Unmarshal are expensive operations.
	b, _ := json.Marshal(q)
	json.Unmarshal(b, &query)

	queryResult := make(map[int]struct{})
	// Do the query
	tiedot.EvalQuery(query, col, &queryResult)
	return flatRes(col, &queryResult)

}

func flatRes(col *tiedot.Col, queryResult *map[int]struct{}) (resp []byte, err error) {
	// Construct array of results
	resultDocs := make(map[string]interface{}, len(*queryResult))
	for docID := range *queryResult {
		doc, _ := col.Read(docID)
		if doc != nil {
			resultDocs[strconv.Itoa(docID)] = doc
		}
	}

	// Serialize the array
	resp, err = json.Marshal(resultDocs)
	if err != nil {
		return nil, err
	}
	return
}
