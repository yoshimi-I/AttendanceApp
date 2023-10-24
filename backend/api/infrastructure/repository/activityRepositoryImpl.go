package repository

import (
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
func (a ActivityRepositoryImpl) FindActivity(id int) error {
	var attendance orm_model.Attendance

	if err := a.db.First(&attendance, id).Error; err != nil {
		return err
	}

	return nil
}
func (a ActivityRepositoryImpl) PostStartActivity(attendance *model.Attendance) (*model.Attendance, error) {
	entity := &orm_model.Attendance{
		UserID:         attendance.UserID,
		AttendanceType: attendance.AttendanceType,
		StartTime:      attendance.StartTime,
		Date:           attendance.Date,
	}

	if err := a.db.Create(entity).Error; err != nil {
		return nil, err
	}
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

func (a ActivityRepositoryImpl) PutStudyActivity(attendance *model.Attendance) error {
	entity := &orm_model.Attendance{
		StartTime: attendance.StartTime,
		EndTime:   attendance.EndTime,
	}
	id := attendance.ID

	if err := a.db.Model(&orm_model.Attendance{}).Where("id = ?", id).Updates(entity).Error; err != nil {
		return err
	}

	return nil
}

func (a ActivityRepositoryImpl) DeleteStudyActivity(userId int, date string) error {

	if err := a.db.Where("user_id = ? AND date = ?", userId, date).Delete(&orm_model.Attendance{}).Error; err != nil {
		return err
	}

	return nil
}
