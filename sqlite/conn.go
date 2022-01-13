package sqlite

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

//"?cache=shared"

const (
	ERR_UNDEFINED_DSN = "connection error: SQLite DSN '%s' is undefined\n"
)

// MustConnect returns a new DB connection
func MustConnect(dsn string, opts ...Option) *sql.DB {
	db, err := connect(dsn, opts...)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

// MustConnect returns a new DB connection
func MustConnectENV(dsn string, opts ...Option) *sql.DB {
	dsn, ok := os.LookupEnv(dsn)
	if !ok {
		log.Fatalf(ERR_UNDEFINED_DSN, dsn)
	}

	db, err := connect(dsn, opts...)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}

// MustConnect returns a new DB connection
func connect(dsn string, opts ...Option) (*sql.DB, error) {
	//"?cache=shared"
	db, err := sql.Open("sqlite3", dsn)
	if err != nil {
		return nil, err
	}

	for _, opt := range opts {
		err = opt(db)
		if err != nil {
			return nil, err
		}
	}

	return db, nil
}
