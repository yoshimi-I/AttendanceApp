package repository

import (
	"fmt"
	"gorm.io/gorm"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	"work-management-app/infrastructure/database/orm"
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
		ID:             attendance.ID,
		UserID:         attendance.UserID,
		AttendanceType: model.IntToActionEnum(attendance.AttendanceType),
		Time:           attendance.Time,
		Year:           attendance.Year,
		Date:           attendance.Date,
	}

	return res, nil
}

// PostActivity 活動を追加する
func (a ActivityRepositoryImpl) PostActivity(attendance *model.Attendance, tx repository.Transaction) (*model.Attendance, error) {

	// 受け取ったトランザクションをormに整形
	conn := ConvertOrm(tx)
	if conn != nil {
		conn = a.db
	}

	entity := &orm_model.Attendance{
		UserID:         attendance.UserID,
		AttendanceType: attendance.AttendanceType.ToInt(),
		Time:           attendance.Time,
		Date:           attendance.Date,
		Year:           attendance.Year,
	}

	if err := conn.Create(entity).Error; err != nil {
		return nil, err
	}

	attendance.ID = entity.ID
	return attendance, nil
}

// PutActivity 活動を編集する
func (a ActivityRepositoryImpl) PutActivity(attendance *model.Attendance, tx repository.Transaction) (*model.Attendance, error) {

	// 受け取ったトランザクションをormに整形
	conn := ConvertOrm(tx)
	if conn != nil {
		conn = a.db
	}

	entity := &orm_model.Attendance{
		Time: attendance.Time,
	}
	id := attendance.ID

	if err := conn.Model(&orm_model.Attendance{}).Where("id = ?", id).Updates(entity).Error; err != nil {
		return nil, err
	}

	return attendance, nil
}

// DeleteActivity　活動を削除する
func (a ActivityRepositoryImpl) DeleteActivity(id int, tx repository.Transaction) error {

	// 受け取ったトランザクションをormに整形
	conn := ConvertOrm(tx)
	if conn != nil {
		conn = a.db
	}
	
	if err := conn.Where("id = ?", id).Delete(&orm_model.Attendance{}).Error; err != nil {
		return err
	}

	return nil
}
