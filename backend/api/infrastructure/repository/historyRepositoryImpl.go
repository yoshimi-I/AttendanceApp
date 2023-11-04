package repository

import (
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	orm_model "work-management-app/infrastructure/orm"

	"gorm.io/gorm"
)

type HistoryRepoImpl struct {
	db *gorm.DB
}

func NewHistoryRepository(db *gorm.DB) repository.HistoryRepository {
	return &HistoryRepoImpl{
		db: db,
	}
}

func (s *HistoryRepoImpl) ReadAllHistory(userID int, year int) ([]model.Attendance, error) {
	var activities []orm_model.Attendance
	var res []model.Attendance
	err := s.db.Where("user_id = ?", userID).Where("year = ?", year).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	for _, activity := range activities {
		resItem := model.Attendance{
			ID:             activity.ID,
			UserID:         activity.UserID,
			AttendanceType: activity.AttendanceType,
			StartTime:      activity.StartTime,
			EndTime:        activity.EndTime,
			Date:           activity.Date,
		}
		res = append(res, resItem)
	}
	return res, nil
}

func (s *HistoryRepoImpl) ReadHistoryByDate(userID int, date string) ([]model.Attendance, error) {
	var activities []orm_model.Attendance
	var res []model.Attendance
	err := s.db.Where("user_id = ?", userID).Where("date = ?", date).Find(&activities).Error
	if err != nil {
		return nil, err
	}
	for _, activity := range activities {
		resItem := model.Attendance{
			ID:             activity.ID,
			UserID:         activity.UserID,
			AttendanceType: activity.AttendanceType,
			StartTime:      activity.StartTime,
			EndTime:        activity.EndTime,
			Date:           activity.Date,
		}
		res = append(res, resItem)
	}
	return res, nil
}
