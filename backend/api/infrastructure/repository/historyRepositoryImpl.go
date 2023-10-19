package repository

import (
	"github.com/yoshimi-I/AttendanceApp/domain/model"
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	orm_model "github.com/yoshimi-I/AttendanceApp/infrastructure/orm"
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

func (s *HistoryRepoImpl) GetAllHistory(userID int) ([]model.Activity, error) {
	var attendances = &orm_model.Attendance{}

	if err := s.db.Joins("AttendanceType").Where("user_id = ?", userID).Find(&attendances).Error; err != nil {
		return nil, err
	}

	return attendances, nil
}

func (s *HistoryRepoImpl) GetHistoryByDate(userID int, date string) (model.Activity, error) {
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
