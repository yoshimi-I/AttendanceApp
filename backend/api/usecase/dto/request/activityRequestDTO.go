package request

import (
	"fmt"
	"time"
)

type ActivityStartRequestDTO struct {
	UserID         int       `json:"user_id"`
	AttendanceType int       `json:"attendance_type"`
	StartTime      time.Time `json:"start_time"`
}

type ActivityEditRequestDTO struct {
	UserID         int       `json:"user_id"`
	AttendanceType int       `json:"attendance_type"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
}

func (a ActivityStartRequestDTO) Year() int {
	return a.StartTime.Year()
}

func (a ActivityStartRequestDTO) Date() string {
	return fmt.Sprintf("%d-%02d-%02d", a.StartTime.Year(), a.StartTime.Month(), a.StartTime.Day())
}

func (a ActivityEditRequestDTO) Year() int {
	return a.StartTime.Year()
}

func (a ActivityEditRequestDTO) Date() string {
	return fmt.Sprintf("%d-%02d-%02d", a.StartTime.Year(), a.StartTime.Month(), a.StartTime.Day())
}

func (a ActivityEditRequestDTO) ShiftTime(year, month, day int) string {
	afTime := a.StartTime.AddDate(year, month, day)
	return fmt.Sprintf("%d-%02d-%02d", afTime.Year(), afTime.Month(), afTime.Day())

}
