package repository

import "github.com/yoshimi-I/AttendanceApp/domain/model"

type HistoryRepository interface {
	GetAllHistory() ([]model.Activities, error)
	GetHistoryByDate(date string) (model.Activities, error)
}
