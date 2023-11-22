package repository

import (
	"gorm.io/gorm"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	"work-management-app/infrastructure/database/orm"
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
			AttendanceType: model.IntToActionEnum(activity.AttendanceType),
			Time:           activity.Time,
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
			AttendanceType: model.IntToActionEnum(activity.AttendanceType),
			Time:           activity.Time,
			Date:           activity.Date,
		}
		res = append(res, resItem)
	}
	return res, nil
}
