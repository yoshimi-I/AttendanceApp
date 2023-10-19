package repository

import "github.com/yoshimi-I/AttendanceApp/domain/model"

type HistoryRepository interface {
	GetAllHistory(userID int) ([]model.Activity, error)
	GetHistoryByDate(userID int, date string) (model.Activity, error)
}
