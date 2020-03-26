package mysql

import (
	"crypto/tls"
	"crypto/x509"
	"io/ioutil"
	"log"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

func TLS(configName, certPath string) {
	if certPath == "" {
		return
	}

	// to get TLS to work you must use the actual DNS hostname of the DB in the DSN string and add that same name to /etc/hosts
	rootCertPool := x509.NewCertPool()
	pem, err := ioutil.ReadFile(certPath)
	if err != nil {
		log.Fatal(err)
	}

	if ok := rootCertPool.AppendCertsFromPEM(pem); !ok {
		log.Fatal("Failed to append PEM.")
	}

	err = mysql.RegisterTLSConfig(configName, &tls.Config{
		RootCAs: rootCertPool,
	})

	if err != nil {
		log.Fatal(err)
	}

}
