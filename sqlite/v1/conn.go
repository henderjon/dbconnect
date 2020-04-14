package sqlite

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

const (
	ERR_UNDEFINED_DSN = "connection error: SQLite DSN '%s' is undefined"
)

// MustConnect returns a new DB connection
func MustConnect(dsn string, maxconns int) (*sql.DB, error) {

	dsn, ok := os.LookupEnv(dsn)
	if ok != true {
		err := fmt.Errorf(ERR_UNDEFINED_DSN, dsn)
		log.Fatal(err)
	}

	db, err := sql.Open("sqlite3", "file:"+dsn+"?cache=shared")
	if err != nil {
		log.Fatal(err)
	}

	db.SetMaxOpenConns(1) // maxconns
	setupQueries := []string{
		`PRAGMA encoding="UTF-8";`,
		`PRAGMA journal_mode=WAL;`,
		`PRAGMA page_size=4096;`,
		`PRAGMA cache_size=-2000000;`,
	}
	for _, qry := range setupQueries {
		_, err = db.Exec(qry)
		if err != nil {
			log.Fatal(fmt.Errorf("%s: %s", err, qry))
		}
	}

	// defer db.Close()
	return db, nil
}

// Vacuum rebuilds the DB file to keep it tidy
func vacuum(db *sql.DB) error {
	sql := `VACUUM;`
	_, err := db.Exec(sql)
	return fmt.Errorf("vacuum failed: %s", err)
}

// WalCheckpoint commits the WAL; DO NOT USE
func walCheckpoint(db *sql.DB) error {
	return nil
	// sql := `PRAGMA wal_checkpoint(TRUNCATE);`
	// _, err := db.Exec(sql)
	// return fmt.Errorf("wal_checkpoint failed: %s", err)
}
