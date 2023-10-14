package repository

import (
	"github.com/yoshimi-I/AttendanceApp/domain/model"
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	"gorm.io/gorm"
	"time"
)

type StudyHistoryImpl struct {
	db *gorm.DB
}

func NewStudyHistoryRepository(db *gorm.DB) repository.StudyHistoryRepository {
	return &StudyHistoryImpl{
		db: db,
	}
}

func (s *StudyHistoryImpl) GetAllHistory() []model.Activities {
	a := model.Activities{
		ActivityDate: time.Time{},
		Notes:        "",
		Plays:        nil,
		Studies:      nil,
		Breaks:       nil,
		SumTime:      2,
	}
	b := model.Activities{
		ActivityDate: time.Time{},
		Notes:        "",
		Plays:        nil,
		Studies:      nil,
		Breaks:       nil,
		SumTime:      0,
	}
	return []model.Activities{a, b}
}

func (s *StudyHistoryImpl) GetHistoryByDate(date string) model.Activities {
	a := model.Activities{
		ActivityDate: time.Time{},
		Notes:        "",
		Plays:        nil,
		Studies:      nil,
		Breaks:       nil,
		SumTime:      0,
	}
	return a
}
