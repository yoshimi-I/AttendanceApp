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
	PostActivity(attendance *model.Attendance, tx Transaction) (*model.Attendance, error)
	PutActivity(attendance *model.Attendance, tx Transaction) (*model.Attendance, error)
	DeleteActivity(id int, tx Transaction) error
}
