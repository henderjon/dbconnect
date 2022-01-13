package mysql

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

// NOTE:
// user:password@tcp(addr:port)/db?args
// ?parseTime=true&loc=UTC

// these should be equal, cause problems if they are not
// db.SetMaxIdleConns(maxconns)
// db.SetMaxOpenConns(maxconns)
// db.SetConnMaxLifetime(0)

const (
	ERR_UNDEFINED_DSN = "connection error: MySQL DSN '%s' is undefined\n"
)

// Connect connects to the given DSN and fails quietly
func Connect(dsn string, tls TLSOption, opts ...Option) *sql.DB {
	db, err := connect(dsn, tls, opts...)
	if err != nil {
		log.Println(err)
	}
	return db
}

// MustConnect connects to the given DSN and fails loudly if unable to do so.
func MustConnect(dsn string, tls TLSOption, opts ...Option) *sql.DB {
	db, err := connect(dsn, tls, opts...)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

// ConnectENV connects to the DSN at the given ENV VAR and fails quietly
func ConnectENV(dsn string, tls TLSOption, opts ...Option) *sql.DB {
	dsn, ok := os.LookupEnv(dsn)
	if !ok {
		log.Fatalf(ERR_UNDEFINED_DSN, dsn)
	}

	db, err := connect(dsn, tls, opts...)
	if err != nil {
		log.Println(err)
	}
	return db
}

// MustConnectENV connects to the DSN at the given ENV VAR and fails loudly if unable to do so.
func MustConnectENV(dsn string, tls TLSOption, opts ...Option) *sql.DB {
	dsn, ok := os.LookupEnv(dsn)
	if !ok {
		log.Fatalf(ERR_UNDEFINED_DSN, dsn)
	}

	db, err := connect(dsn, tls, opts...)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func connect(dsn string, tls TLSOption, opts ...Option) (*sql.DB, error) {
	if tls != nil {
		err := tls()
		if err != nil {
			return nil, err
		}
	}

	db, err := sql.Open("mysql", dsn)
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
