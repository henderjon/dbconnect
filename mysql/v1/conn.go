package mysql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

const (
	ERR_UNDEFINED_DSN = "connection error: MySQL DSN '%s' is undefined\n"
)

// Connect connects to the DSN at the given ENV VAR and fails quietly
func Connect(dsn string, tls TLSOption, opts ...Option) *sql.DB {

	// NOTE:
	// user:password@tcp(addr:port)/db?args
	// ?parseTime=true&loc=UTC
	dsn, ok := os.LookupEnv(dsn)
	if ok != true {
		log.Printf(ERR_UNDEFINED_DSN, dsn)
		return nil
	}

	if tls != nil {
		err := tls()
		if err != nil {
			log.Println(err)
		}
	}

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Println(err)
		return nil
	}

	for _, opt := range opts {
		err = opt(db)
		if err != nil {
			log.Println(err)
		}
	}

	// these should be equal, cause problems if they are not
	// db.SetMaxIdleConns(maxconns)
	// db.SetMaxOpenConns(maxconns)
	// db.SetConnMaxLifetime(0)
	return db

}

// MustConnect connects to the DSN at the given ENV VAR and fails loudly if unable to do so.
func MustConnect(dsn string, tls TLSOption, opts ...Option) *sql.DB {

	// NOTE:
	// user:password@tcp(addr:port)/db?args
	// ?parseTime=true&loc=UTC
	dsn, ok := os.LookupEnv(dsn)
	if ok != true {
		log.Fatalf(ERR_UNDEFINED_DSN, dsn)
	}

	if tls != nil {
		err := tls()
		if err != nil {
			log.Fatalln(err)
		}
	}

	db, err := sql.Open("mysql", dsn)
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
