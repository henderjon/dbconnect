package sqlite

import (
	"database/sql"
	"fmt"
)

// https://www.sqlite.org/pragma.html
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

// WalCheckpoint checkpoints the Write-Ahead-Log
func WALCheckpoint(mode string) Option {
	return Option(func(db *sql.DB) error {
		qry := fmt.Sprintf(`PRAGMA wal_checkpoint(%s);`, mode)
		_, err := db.Exec(qry)
		if err != nil {
			return fmt.Errorf("%s: %s", err, qry)
		}
		return nil
	})
}

// MaxOpenConnections sets the max open connections. Since the SQL package
// handles the connection pool, it's best to set this to 1 (one) so that the
// SQL package essentially wraps all writes in a mutex. This will slow down
// writes but will help avoid a locked DB.
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

// JournalMode sets the journal mode for the DB, typically WAL
func JournalMode(mode string) Option {
	return Option(func(db *sql.DB) error {
		qry := fmt.Sprintf(`PRAGMA journal_mode=%s;`, mode)
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

// CacheSize defaults to '-2000' kibibytes which is 2.048MB
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
