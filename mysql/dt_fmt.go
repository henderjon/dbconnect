package mysql

import (
	"time"

	"github.com/henderjon/dbconnect/datetimes"
)

const ( // Mon Jan 2 15:04:05 MST 2006
	FmtTimestampDay    = datetimes.FmtTimestampDay
	FmtTimestampHour   = datetimes.FmtTimestampHour
	FmtTimestampMinute = datetimes.FmtTimestampMinute
	FmtDateTime        = datetimes.FmtDateTime
	FmtDate            = datetimes.FmtDate
	FmtTime            = datetimes.FmtTime

	// mysql datetime format strings
	MySQLFmtDatetime = "%Y-%m-%d %H:%i:%s"
	MySQLFmtMinute   = "%Y-%m-%d %H:%i"
	MySQLFmtHour     = "%Y-%m-%d %H"
	MySQLFmtDay      = "%Y-%m-%d"
)

func ParseMySQLDateTime(t string) time.Time {
	timestamp, err := time.Parse(FmtDateTime, t)
	if err != nil {
		panic(err)
	}
	return timestamp
}
