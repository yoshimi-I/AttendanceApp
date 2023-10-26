package repository

import "github.com/yoshimi-I/AttendanceApp/domain/model"

type HistoryRepository interface {
	ReadAllHistory(userID int, year int) ([]model.Attendance, error)
	ReadHistoryByDate(userID int, date string) ([]model.Attendance, error)
}
