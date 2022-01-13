package mysql

import (
	"crypto/tls"
	"crypto/x509"
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type TLSOption func() error

type Option func(db *sql.DB) error

func TLS(configName, certPath string) TLSOption {
	return TLSOption(func() error {
		if certPath == "" {
			return fmt.Errorf("Failed to open PEM: '%s'", certPath)
		}

		// to get TLS to work you must use the actual DNS hostname of the DB in the DSN string and add that same name to /etc/hosts
		rootCertPool := x509.NewCertPool()
		pem, err := ioutil.ReadFile(certPath)
		if err != nil {
			return err
		}

		if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
			return fmt.Errorf("Failed to append PEM: '%s'", certPath)
		}

		err = mysql.RegisterTLSConfig(configName, &tls.Config{
			RootCAs: rootCertPool,
		})

		return err
	})
}

func MaxIdleConnections(conns int) Option {
	return Option(func(db *sql.DB) error {
		db.SetMaxIdleConns(conns)
		return nil
	})
}

func MaxOpenConnections(conns int) Option {
	return Option(func(db *sql.DB) error {
		db.SetMaxOpenConns(conns)
		return nil
	})
}

func ConnMaxLifetime(d time.Duration) Option {
	return Option(func(db *sql.DB) error {
		db.SetConnMaxLifetime(d)
		return nil
	})
}
