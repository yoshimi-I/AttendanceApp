package repository

import (
	"github.com/yoshimi-I/AttendanceApp/domain/model"
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	"gorm.io/gorm"
	"time"
)

type HistoryRepoImpl struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) repository.HistoryRepository {
	return &HistoryRepoImpl{
		db: db,
	}
}

func (s *HistoryRepoImpl) GetAllHistory() ([]model.Activities, error) {
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
	return []model.Activities{a, b}, nil
}

func (s *HistoryRepoImpl) GetHistoryByDate(date string) (model.Activities, error) {
	a := model.Activities{
		ActivityDate: time.Time{},
		Notes:        "",
		Plays:        nil,
		Studies:      nil,
		Breaks:       nil,
		SumTime:      0,
	}
	return a, nil
}
