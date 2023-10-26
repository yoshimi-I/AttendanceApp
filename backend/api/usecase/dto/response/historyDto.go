package response

import (
	"time"
)

type AllHistoryDto struct {
	Day      string `json:"day"`
	Activity int    `json:"activity"`
}

type ActivityDetail struct {
	Type      string    `json:"type"`
	StartTime time.Time `json:"startTime"`
	EndTime	 time.Time `json:"endTime"`
	Timestamp time.Time `json:"timestamp"`
}

type HistoryByDateDto struct {
	Date       string           `json:"date"`
	Activities []ActivityDetail `json:"activities"`
}

func ConvertAllHistory(day string, activity int) *AllHistoryDto {
	return &AllHistoryDto{
		Day:      day,
		Activity: activity,
	}
}
