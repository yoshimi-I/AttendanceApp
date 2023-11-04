package repository

import (
	"time"
	"work-management-app/domain/model"
)

type ActivityInput struct {
	AttendanceType int
	CurrentTime    time.Time
	Date           string
}

type ActivityRepository interface {
	FindActivity(id int) (*model.Attendance, error)
	PostStartActivity(attendance *model.Attendance) (*model.Attendance, error)
	PostEndActivity(attendance *model.Attendance) (*model.Attendance, error)
	PutActivity(attendance *model.Attendance) (*model.Attendance, error)
	DeleteActivity(id int) error
}
