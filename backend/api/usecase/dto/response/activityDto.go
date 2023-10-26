package response

import "time"

type ActivityResponseDTO struct {
	ID             int       `json:"id"`
	UserID         int       `json:"userID"`
	AttendanceType string    `json:"attendanceType"`
	StartTime      time.Time `json:"startTime"`
	EndTime        time.Time `json:"endTime"`
	Year           int       `json:"year"`
	Date           string    `json:"date"`
}
