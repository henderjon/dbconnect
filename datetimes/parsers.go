package datetimes

const ( // Mon Jan 2 15:04:05 MST 2006
	// go time format strings
	FmtDateTime        = "2006-01-02 15:04:05"
	FmtDate            = "2006-01-02"
	FmtTime            = "15:04:05"
	FmtTimestampDay    = "2006-01-02 00:00:00"
	FmtTimestampHour   = "2006-01-02 15:00:00"
	FmtTimestampMinute = "2006-01-02 15:04:00"

	// PrefixFormatMinutes     = "200601021504"     // time.Now().UTC().Format(...) month-day-hour-min // Mon Jan 2 15:04:05 -0700 MST 2006
	// PrefixFormatHours       = "2006010215"       // time.Now().UTC().Format(...) month-day-hour     // Mon Jan 2 15:04:05 -0700 MST 2006
	// PrefixFormatDays        = "20060102"         // time.Now().UTC().Format(...) month-day          // Mon Jan 2 15:04:05 -0700 MST 2006
	// FilePrefixFormatMinutes = "2006/01/02/15/04" // time.Now().UTC().Format(...) month-day-hour-min // Mon Jan 2 15:04:05 -0700 MST 2006
	// FilePrefixFormatHours   = "2006/01/02/15"    // time.Now().UTC().Format(...) month-day-hour     // Mon Jan 2 15:04:05 -0700 MST 2006
	// FilePrefixFormatDays    = "2006/01/02/"      // time.Now().UTC().Format(...) month-day          // Mon Jan 2 15:04:05 -0700 MST 2006

)
