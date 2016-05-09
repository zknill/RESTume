package service

import (
	"fmt"
	"os"

	tiedot "github.com/HouzuoGuo/tiedot/db"
	log "github.com/golang/glog"
)

// Resource represents an external resource such as a database
type Resource interface {
	Init()
}

// Database represents the tiedot NoSQL datastore.
type Database struct {
	Name     string
	Location string
	Data     *tiedot.DB
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
	os.RemoveAll(db.Location)

	data, dbErr := tiedot.OpenDB(db.Location)
	if dbErr != nil {
		log.Error(dbConnectionError{
			msg: "Failed to connect to the tiedot database",
			err: dbErr,
		})
	}
	db.Data = data
}
