package dto

import "time"

type ALlHistoryDto struct {
	Day      string `json:"day"`
	Activity int    `json:"activity"`
}

type ActivityDetail struct {
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}

type HistoryByDateDto struct {
	Date       string           `json:"date"`
	Activities []ActivityDetail `json:"activities"`
}
