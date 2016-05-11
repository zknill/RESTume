package resources

import (
	"encoding/json"
	"fmt"
	"net/http"

	tiedot "github.com/HouzuoGuo/tiedot/db"
	log "github.com/golang/glog"
)

// Database represents the tiedot NoSQL datastore.
type Database struct {
	Name        string
	Location    string
	Collections []*Collection
	Data        *tiedot.DB
}

type Collection struct {
	Col   string
	Index []string
}

type dbConnectionError struct {
	msg string
	err error
}

func (e *dbConnectionError) Error() string {
	return fmt.Sprintf("%s, %s", e.msg, e.err)
}

// Init does the setup for the database. It also inherently implements Resource
func (db *Database) Init() {
	data, dbErr := tiedot.OpenDB(db.Location)
	if dbErr != nil {
		log.Error(dbConnectionError{
			msg: "Failed to connect to the tiedot database",
			err: dbErr,
		})
	}

	// Set up the collections - throw away the error for now.
	for _, c := range db.Collections {
		data.Create(c.Col)
		data.Use(c.Col).Index(c.Index)
	}

	db.Data = data
}

func DBInsert(r *http.Request, col *tiedot.Col) (id int, err error) {
	data := map[string]interface{}{}
	decoder := json.NewDecoder(r.Body)
	decoder.Decode(&data)

	id, err = col.Insert(data)
	return
}
