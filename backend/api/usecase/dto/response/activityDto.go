package response

import (
	"log"
	"time"
)

type ActivityResponseDTO struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	AttendanceType string    `json:"attendance_type"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	Year           int       `json:"year"`
	Date           string    `json:"date"`
}

// ActivityTimeResponseDTO アクティビティのレスポンスデータ
type ActivityTimeResponseDTO struct {
	Date    string
	SumTime int
}

func ConvertActivityTime(actionType int) string {
	var Type string
	switch actionType {
	case 1:
		Type = "作業"
	case 2:
		Type = "休憩"
	case 3:
		Type = "その他"
	default:
		log.Printf("Invalid AttendanceType: %d", actionType)
	}
	return Type
}
