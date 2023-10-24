package repository

import (
	"github.com/yoshimi-I/AttendanceApp/domain/model"
	"time"
)

type ActivityInput struct {
	AttendanceType int
	CurrentTime    time.Time
	Date           string
}

type ActivityRepository interface {
	FindActivity(id int) error
	PostStartActivity(attendance *model.Attendance) (*model.Attendance, error)
	PostEndActivity(attendance *model.Attendance) (*model.Attendance, error)
	PutStudyActivity(attendance *model.Attendance) error
	DeleteStudyActivity(userId int, data string) error
}
