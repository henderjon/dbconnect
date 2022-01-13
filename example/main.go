package main

import (
	"fmt"

	mysql "github.com/henderjon/dbconnect/mysql"
	sqlite "github.com/henderjon/dbconnect/sqlite"
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
		sqlite.Encoding(`UTF-8`),
		sqlite.JournalMode(`WAL`),
		sqlite.PageSize(4096),
		sqlite.CacheSize(-200000),
	)
	fmt.Println(db.Ping())
}
