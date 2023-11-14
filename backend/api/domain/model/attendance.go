package model

import "time"

type Attendance struct {
	ID             int
	UserId         int
	AttendanceType ActionEnum
	Time           time.Time
	Year           int
	Date           string
}
