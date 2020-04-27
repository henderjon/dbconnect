package sqlite

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Option func(db *sql.DB) error

// Vacuum rebuilds the DB file to keep it tidy
func Vacuum() Option {
	return Option(func(db *sql.DB) error {
		sql := `VACUUM;`
		_, err := db.Exec(sql)
		if err != nil {
			return fmt.Errorf("vacuum failed: %s", err)
		}
		return nil
	})
}

// WalCheckpoint commits the WAL; DO NOT USE
func WALCheckpoint() Option {
	return Option(func(db *sql.DB) error {
		sql := `PRAGMA wal_checkpoint(TRUNCATE);`
		_, err := db.Exec(sql)
		if err != nil {
			return fmt.Errorf("wal_checkpoint failed: %s", err)
		}
		return nil
	})
}

func MaxOpenConnections(conns int) Option {
	return Option(func(db *sql.DB) error {
		db.SetMaxOpenConns(conns)
		return nil
	})
}

// Encoding set the encoding of the DB; should be "UTF-8"
func Encoding(charset string) Option {
	return Option(func(db *sql.DB) error {
		qry := fmt.Sprintf(`PRAGMA encoding='%s';`, charset)
		_, err := db.Exec(qry)
		if err != nil {
			return fmt.Errorf("%s: %s", err, qry)
		}
		return nil
	})
}

func EncodingUTF8() Option {
	return Option(func(db *sql.DB) error {
		qry := `PRAGMA encoding='UTF-8';`
		_, err := db.Exec(qry)
		if err != nil {
			return fmt.Errorf("%s: %s", err, qry)
		}
		return nil
	})
}

func Journal(mode string) Option {
	return Option(func(db *sql.DB) error {
		qry := fmt.Sprintf(`PRAGMA journal_mode=%s;`, mode)
		_, err := db.Exec(qry)
		if err != nil {
			return fmt.Errorf("%s: %s", err, qry)
		}
		return nil
	})
}

func JournalWAL() Option {
	return Option(func(db *sql.DB) error {
		qry := `PRAGMA journal_mode=WAL;`
		_, err := db.Exec(qry)
		if err != nil {
			return fmt.Errorf("%s: %s", err, qry)
		}
		return nil
	})
}

// PageSize the default is 4096 and is recommended
func PageSize(pageSize int) Option {
	return Option(func(db *sql.DB) error {
		qry := fmt.Sprintf(`PRAGMA page_size=%d;`, pageSize)
		_, err := db.Exec(qry)
		if err != nil {
			return fmt.Errorf("%s: %s", err, qry)
		}
		return nil
	})
}

// CacheSize defaults to '-2000' which is 2.048MB
func CacheSize(cacheSize int) Option {
	return Option(func(db *sql.DB) error {
		qry := fmt.Sprintf(`PRAGMA cache_size=%d;`, cacheSize)
		_, err := db.Exec(qry)
		if err != nil {
			return fmt.Errorf("%s: %s", err, qry)
		}
		return nil
	})
}

// BusyTimeout 1000 milliseconds == 1 second
func BusyTimeout(milliseconds int) Option {
	return Option(func(db *sql.DB) error {
		qry := fmt.Sprintf(`PRAGMA busy_timeout=%d;`, milliseconds)
		_, err := db.Exec(qry)
		if err != nil {
			return fmt.Errorf("%s: %s", err, qry)
		}
		return nil
	})
}
