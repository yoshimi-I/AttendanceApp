package repository

import (
	"fmt"
	"github.com/yoshimi-I/AttendanceApp/domain/model"
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	orm_model "github.com/yoshimi-I/AttendanceApp/infrastructure/orm"
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

	result := &model.Attendance{
		ID:             attendance.ID,
		UserID:         attendance.UserID,
		AttendanceType: attendance.AttendanceType,
		StartTime:      attendance.StartTime,
		EndTime:        attendance.EndTime,
		Year:           attendance.Year,
		Date:           attendance.Date,
	}

	return result, nil
}

func (a ActivityRepositoryImpl) PostStartActivity(attendance *model.Attendance) (*model.Attendance, error) {
	entity := &orm_model.Attendance{
		UserID:         attendance.UserID,
		AttendanceType: attendance.AttendanceType,
		StartTime:      attendance.StartTime,
		EndTime:        attendance.EndTime,
		Date:           attendance.Date,
		Year:           attendance.Year,
	}

	if err := a.db.Create(entity).Error; err != nil {
		return nil, err
	}

	attendance.ID = entity.ID
	return attendance, nil
}

func (a ActivityRepositoryImpl) PostEndActivity(attendance *model.Attendance) (*model.Attendance, error) {
	id := attendance.ID

	entity := &orm_model.Attendance{
		EndTime: attendance.EndTime,
	}
	if err := a.db.Model(&orm_model.Attendance{}).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}

	return attendance, nil
}

func (a ActivityRepositoryImpl) PutActivity(attendance *model.Attendance) (*model.Attendance, error) {
	entity := &orm_model.Attendance{
		StartTime: attendance.StartTime,
		EndTime:   attendance.EndTime,
	}
	id := attendance.ID

	if err := a.db.Model(&orm_model.Attendance{}).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}

	return attendance, nil
}

func (a ActivityRepositoryImpl) DeleteActivity(id int) error {

	if err := a.db.Where("id = ?", id).Delete(&orm_model.Attendance{}).Error; err != nil {
		return err
	}

	return nil
}
