package repository

import (
	"fmt"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	orm_model "work-management-app/infrastructure/orm"

	"gorm.io/gorm"
)

type ActivityRepositoryImpl struct {
	db *gorm.DB
}

func NewActivityRepository(db *gorm.DB) repository.ActivityRepository {
	return &ActivityRepositoryImpl{
		db: db,
	}
}
func (a ActivityRepositoryImpl) FindActivity(id int) (*model.Attendance, error) {
	var attendance orm_model.Attendance

	if err := a.db.Where("id = ?", id).First(&attendance).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no record found with id: %d", id)
		}
		return nil, err
	}

	res := &model.Attendance{
		ID:             attendance.Id,
		UserId:         attendance.UserId,
		AttendanceType: model.IntToActionEnum(attendance.AttendanceType),
		Time:           attendance.Time,
		Year:           attendance.Year,
		Date:           attendance.Date,
	}

	return res, nil
}

// PostActivity 活動を追加する
func (a ActivityRepositoryImpl) PostActivity(attendance *model.Attendance) (*model.Attendance, error) {
	entity := &orm_model.Attendance{
		UserId:         attendance.UserId,
		AttendanceType: attendance.AttendanceType.ToInt(),
		Time:           attendance.Time,
		Date:           attendance.Date,
		Year:           attendance.Year,
	}

	if err := a.db.Create(entity).Error; err != nil {
		return nil, err
	}

	attendance.ID = entity.Id
	return attendance, nil
}

// PutActivity 活動を編集する
func (a ActivityRepositoryImpl) PutActivity(attendance *model.Attendance) (*model.Attendance, error) {
	entity := &orm_model.Attendance{
		Time: attendance.Time,
	}
	id := attendance.ID

	if err := a.db.Model(&orm_model.Attendance{}).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}

	return attendance, nil
}

// DeleteActivity　活動を削除する
func (a ActivityRepositoryImpl) DeleteActivity(id int) error {
	if err := a.db.Where("id = ?", id).Delete(&orm_model.Attendance{}).Error; err != nil {
		return err
	}

	return nil
}
