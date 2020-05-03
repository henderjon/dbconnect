package sqlite

import (
	"database/sql"
	"log"
	"os"

	_ "rsc.io/sqlite"
)

const (
	ERR_UNDEFINED_DSN = "connection error: SQLite DSN '%s' is undefined\n"
)

// MustConnect returns a new DB connection
func MustConnect(dsn string, opts ...Option) *sql.DB {

	dsn, ok := os.LookupEnv(dsn)
	if ok != true {
		log.Fatalf(ERR_UNDEFINED_DSN, dsn)
	}

	//"?cache=shared"
	db, err := sql.Open("sqlite3", "file:"+dsn)
	if err != nil {
		log.Fatalln(err)
	}

	for _, opt := range opts {
		err = opt(db)
		if err != nil {
			log.Fatalln(err)
		}
	}

	return db
}
