package sqlite

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

	// SQLite datetime format strings
	SQLiteFmtDatetime = "%Y-%m-%d %H:%M:%S"
	SQLiteFmtMinute   = "%Y-%m-%d %H:%M"
	SQLiteFmtHour     = "%Y-%m-%d %H"
	SQLiteFmtDay      = "%Y-%m-%d"
)

func ParseSQLiteDateTime(t string) time.Time {
	timestamp, err := time.Parse(FmtDateTime, t)
	if err != nil {
		panic(err)
	}
	return timestamp
}
