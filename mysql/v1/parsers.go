package mysql

import (
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const ( // Mon Jan 2 15:04:05 MST 2006
	// go time format strings
	MySQLTimestampDay    = "2006-01-02"
	MySQLTimestampHour   = "2006-01-02 15:00:00"
	MySQLTimestampMinute = "2006-01-02 15:04:00"
	MySQLDateTime        = "2006-01-02 15:04:05"
	MySQLDate            = "2006-01-02"
	MySQLTime            = "15:04:05"

	// PrefixFormatMinutes     = "200601021504"     // time.Now().UTC().Format(...) month-day-hour-min // Mon Jan 2 15:04:05 -0700 MST 2006
	// PrefixFormatHours       = "2006010215"       // time.Now().UTC().Format(...) month-day-hour     // Mon Jan 2 15:04:05 -0700 MST 2006
	// PrefixFormatDays        = "20060102"         // time.Now().UTC().Format(...) month-day          // Mon Jan 2 15:04:05 -0700 MST 2006
	// FilePrefixFormatMinutes = "2006/01/02/15/04" // time.Now().UTC().Format(...) month-day-hour-min // Mon Jan 2 15:04:05 -0700 MST 2006
	// FilePrefixFormatHours   = "2006/01/02/15"    // time.Now().UTC().Format(...) month-day-hour     // Mon Jan 2 15:04:05 -0700 MST 2006
	// FilePrefixFormatDays    = "2006/01/02/"      // time.Now().UTC().Format(...) month-day          // Mon Jan 2 15:04:05 -0700 MST 2006

	// mysql datetime format strings
	MySQLInternalFormat = "%Y-%m-%d %H:%i:%s"
	MySQLDTMinute       = "%Y-%m-%d %H:%i"
	MySQLDTHour         = "%Y-%m-%d %H"
	MySQLDTDay          = "%Y-%m-%d"
)

func ParseMySQLDateTime(t string) time.Time {
	timestamp, err := time.Parse(MySQLDateTime, t)
	if err != nil {
		panic(err)
	}
	return timestamp
}
