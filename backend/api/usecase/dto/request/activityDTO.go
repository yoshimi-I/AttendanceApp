package request

import (
	"fmt"
	"time"
)

type ActivityRequestDTO struct {
	UserKey string    `json:"user_key"`
	Time    time.Time `json:"time"`
}

type ActivityEditRequestDTO struct {
	UserKey        string    `json:"user_key"`
	AttendanceType int       `json:"attendance_type"`
	Time           time.Time `json:"time"`
}

type ActivityDeleteRequestDTO struct {
	UserKey string `json:"user_key"`
}

// Year 年を返す (2023)
func (a ActivityRequestDTO) Year() int {
	return a.Time.Year()
}

// Date 日程を返す ("2023-12-25")
func (a ActivityRequestDTO) Date() string {
	return fmt.Sprintf("%d-%02d-%02d", a.Time.Year(), a.Time.Month(), a.Time.Day())
}

// ShiftTime 移動先の日程を返す  ("2023-12-28")
func (a ActivityRequestDTO) ShiftTime(year, month, day int) string {
	afTime := a.Time.AddDate(year, month, day)
	return fmt.Sprintf("%d-%02d-%02d", afTime.Year(), afTime.Month(), afTime.Day())
}

// Year 年を返す (2023)
func (a ActivityEditRequestDTO) Year() int {
	return a.Time.Year()
}

// Date 日程を返す ("2023-12-25")
func (a ActivityEditRequestDTO) Date() string {
	return fmt.Sprintf("%d-%02d-%02d", a.Time.Year(), a.Time.Month(), a.Time.Day())
}
