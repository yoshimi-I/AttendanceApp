package model

import "time"

type Attendance struct {
	ID             int
	UserID         int
	AttendanceType int
	StartTime      time.Time
	EndTime        time.Time
	Date           string
}
