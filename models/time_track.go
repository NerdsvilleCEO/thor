package models

import "time"

type TimeLog []TimeLogEntry

type TimeLogEntry struct {
	date time.Time
	late bool
	out  bool
}
