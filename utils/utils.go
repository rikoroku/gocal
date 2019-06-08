package utils

import "time"

func StartTimeWithFormatRFC3339() string {
	t := time.Now()
	timezone, _ := time.LoadLocation("Local")
	st := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, timezone)
	return st.Format(time.RFC3339)
}

func EndTimeWithFormatRFC3339() string {
	t := time.Now()
	timezone, _ := time.LoadLocation("Local")
	et := time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 59, timezone)
	return et.Format(time.RFC3339)
}
