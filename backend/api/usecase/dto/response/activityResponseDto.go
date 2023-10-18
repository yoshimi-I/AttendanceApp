package response

import "time"

// ActivityRequestDTO アクティビティのリクエストデータ
type ActivityRequestDTO struct {
	Type      string    `json:"type"`
	Timestamp time.Time `json:"timestamp"`
}

// ActivityResponseDTO アクティビティのレスポンスデータ
type ActivityResponseDTO struct {
	Message      string `json:"message"`
	AttendanceID int    `json:"attendanceId"`
}
