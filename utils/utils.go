package utils

import "time"

func AM00_00() time.Time {
	t := time.Now()
	timezone, _ := time.LoadLocation("Local")
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, timezone)
}

func FM23_59() time.Time {
	t := time.Now()
	timezone, _ := time.LoadLocation("Local")
	return time.Date(t.Year(), t.Month(), t.Day(), 23, 59, 59, 59, timezone)
}
