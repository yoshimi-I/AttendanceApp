package request

import (
	"fmt"
	"time"
)

type ActivityRequestDTO struct {
	UserID         int       `json:"userID"`
	AttendanceType int       `json:"attendanceType"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
	Year           int       `json:"year"`
	Date           time.Time `json:"date"`
}

func ToString(Date time.Time) string {
	return fmt.Sprintf("%04d-%02d-%02d", Date.Year(), Date.Month(), Date.Day())
}
