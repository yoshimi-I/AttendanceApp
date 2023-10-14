package repository

import "github.com/yoshimi-I/AttendanceApp/domain/model"

type StudyHistoryRepository interface {
	GetAllHistory() []model.Activities
	GetHistoryByDate(date string) model.Activities
}
