package request

import (
	"fmt"
	"time"
)

type ActivityRequestDTO struct {
	UserKey string `json:"user_key"`
}

type ActivityEditRequestDTO struct {
	ActivityId int       `json:"activity_id"`
	UserKey    string    `json:"user_key"`
	Time       time.Time `json:"time"`
}

type ActivityDeleteRequestDTO struct {
	ActivityId int    `json:"activity_id"`
	UserKey    string `json:"user_key"`
}

// Year 年を返す (2023)
func (a ActivityEditRequestDTO) Year() int {
	return a.Time.Year()
}

// Date 日程を返す ("2023-12-25")
func (a ActivityEditRequestDTO) Date() string {
	return fmt.Sprintf("%d-%02d-%02d", a.Time.Year(), a.Time.Month(), a.Time.Day())
}
