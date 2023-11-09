package response

import (
	"log"
	"time"
	"work-management-app/domain/model"
)

type ActivityResponseDTO struct {
	ID             int       `json:"id"`
	UserID         int       `json:"user_id"`
	AttendanceType string    `json:"attendance_type"`
	StartTime      time.Time `json:"start_time"`
	EndTime        time.Time `json:"end_time"`
	Year           int       `json:"year"`
	Date           string    `json:"date"`
	Status         string    `json:"status"`
}

// ActivityTimeResponseDTO アクティビティのレスポンスデータ
type ActivityTimeResponseDTO struct {
	Date         string `json:"day"` // フロントに合わせて修正
	ActivityTime int    `json:"activity_time"`
}

func ConvertActivityTime(actionType model.ActionEnum) string {
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
