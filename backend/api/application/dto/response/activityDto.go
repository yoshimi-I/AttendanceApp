package response

import (
	"log"
	"time"
	"work-management-app/domain/model"
)

type ActivityResponseDTO struct {
	Id             int       `json:"id"`
	AttendanceType string    `json:"attendance_type"`
	Time           time.Time `json:"time"`
	Year           int       `json:"year"`
	Date           string    `json:"date"`
	Status         string    `json:"status"`
}

// ActivityTimeResponseDTO アクティビティのレスポンスデータ
type ActivityTimeResponseDTO struct {
	Year         int    `json:"year"`
	Date         string `json:"day"` // フロントに合わせて修正
	ActivityTime int    `json:"activity_time"`
}

func ConvertActivityTime(actionType model.ActionEnum) string {
	var Type string
	switch actionType {
	case 1:
		Type = "作業開始"
	case 2:
		Type = "作業終了"
	case 3:
		Type = "休憩開始"
	case 4:
		Type = "休憩終了"
	case 5:
		Type = "お祈り"
	default:
		log.Printf("Invalid AttendanceType: %d", actionType)
		return "undefined"
	}
	return Type
}
