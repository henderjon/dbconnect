package main

import (
	"fmt"

	"github.com/henderjon/dbconnect/mysql/v1"
	"github.com/henderjon/dbconnect/sqlite/v1"
)

func main() {
	mysqlPing()
	sqlitePing()
}

func mysqlPing() {
	// mysql; assumes an ENV VAR named `GO_TEST_MYSQL_DSN`
	db := mysql.Connect("GO_TEST_MYSQL_DSN", nil,
		mysql.MaxIdleConnections(35),
		mysql.MaxOpenConnections(35),
		mysql.ConnMaxLifetime(0),
	)
	fmt.Println(db.Ping())
}

func sqlitePing() {
	// sqlite; assumes an ENV VAR named `GO_TEST_SQLITE_DSN`
	db := sqlite.MustConnect("GO_TEST_SQLITE_DSN",
		sqlite.MaxOpenConnections(1),
		sqlite.EncodingUTF8(),
		sqlite.JournalWAL(),
		sqlite.PageSize(4096),
		sqlite.CacheSize(-200000),
	)
	fmt.Println(db.Ping())
}
