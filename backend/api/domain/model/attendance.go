package model

import "time"

type Attendance struct {
	ID             int
	UserID         int
	AttendanceType ActionEnum
	StartTime      time.Time
	EndTime        time.Time
	Year           int
	Date           string
}
