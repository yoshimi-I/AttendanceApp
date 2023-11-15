package model

import "time"

type Attendance struct {
	Id             int
	UserId         int
	AttendanceType ActionEnum
	Time           time.Time
	Year           int
	Date           string
}
