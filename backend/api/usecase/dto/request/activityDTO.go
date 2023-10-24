package request

import "time"

type ActivityInput struct {
	UserID         int
	AttendanceType int
	CurrentTime    time.Time
	Date           string
}
