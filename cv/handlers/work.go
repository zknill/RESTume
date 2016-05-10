package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	tiedot "github.com/HouzuoGuo/tiedot/db"
	"github.com/davecgh/go-spew/spew"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
	s "github.com/zknill/RESTume/service"
)

//TODO: refactor this file, there shouldn't be references to the database used. Abstract these.
//TODO: refactor endpoint method, it's too long. split out logically.
//TODO: add authentication to this endpoint. Anyone can add data into the db.

// Work is the work endpoint handler
func Work(w http.ResponseWriter, r *http.Request) error {
	// TODO: architect the context and resources better - this is horrible.
	sh := context.Get(r, s.ContextKey).(*s.ServiceHandler)
	db := sh.Resources["db"].(*s.Database).Data
	col := db.Use("career")

	if r.Method == "POST" {
		if _, err := dbPOST(r, col); err != nil {
			return fmt.Errorf(fmt.Sprint("POST error: %s", err))
		}
		w.Write([]byte(`{"success": "True"}`))
		return nil
	}

	vars := mux.Vars(r)

	spew.Dump(vars)
	comp := vars["company"]

	q := map[string]interface{}{
		"eq": comp,
		"in": []string{"Company"},
	}

	//TODO: make this not shit! marshall and unmarshall are expensive operations.
	b, _ := json.Marshal(q)

	var query interface{}
	queryResult := make(map[int]struct{})

	if err := json.Unmarshal(b, &query); err != nil {
		return err
	}

	spew.Dump(query)

	tiedot.EvalQuery(query, col, &queryResult)

	// Construct array of result
	resultDocs := make(map[string]interface{}, len(queryResult))
	counter := 0
	for docID := range queryResult {
		doc, _ := col.Read(docID)
		if doc != nil {
			resultDocs[strconv.Itoa(docID)] = doc
			counter++
		}
	}
	// Serialize the array
	resp, err := json.Marshal(resultDocs)
	if err != nil {
		return err
	}

	w.Write(resp)

	return nil
}

func dbPOST(r *http.Request, col *tiedot.Col) (id int, err error) {
	data := map[string]interface{}{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	id, err = col.Insert(data)
	return
}
