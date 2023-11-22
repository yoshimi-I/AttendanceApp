package model

import "time"

type Attendance struct {
	ID             int
	UserID         int
	AttendanceType ActionEnum
	Time           time.Time
	Year           int
	Date           string
}
