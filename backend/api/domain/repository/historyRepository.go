package repository

import "work-management-app/domain/model"

type HistoryRepository interface {
	ReadAllHistory(userId int, year int) ([]model.Attendance, error)
	ReadHistoryByDate(userId int, date string) ([]model.Attendance, error)
}
