package repository

import (
	"time"
)

type Activity struct {
	Type      string
	Timestamp time.Time
}

type ActivtyRepository interface {
	PostStudyActivity(activity Activity) error
	PutStudyActivity(id int, activity Activity) error
	DeleteStudyActivity(id int) error
}
