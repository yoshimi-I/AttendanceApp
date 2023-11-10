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
func (a ActivityRepositoryImpl) PostActivity(attendance *model.Attendance) (*model.Attendance, error) {
	entity := &orm_model.Attendance{
		UserID:         attendance.UserID,
		AttendanceType: attendance.AttendanceType.ToInt(),
		Time:           attendance.Time,
		Date:           attendance.Date,
		Year:           attendance.Year,
	}

	if err := a.db.Create(entity).Error; err != nil {
		return nil, err
	}

	attendance.ID = entity.ID
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

// FindUserStatus 現在のユーザーの状態を確認
func (a ActivityRepositoryImpl) FindUserStatus(userID int) (*model.UserStatus, error) {
	var status orm_model.UserStatus

	if err := a.db.Where("user_id = ?", userID).First(&status).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no record found with UserID: %d", userID)
		}
		return nil, err
	}

	res := &model.UserStatus{
		UserID:   status.UserID,
		StatusID: model.IntToStatusEnum(status.StatusID),
	}

	return res, nil

}

// PostUserStatus ユーザーの状態を新規登録
func (a ActivityRepositoryImpl) PostUserStatus(status *model.UserStatus) (*model.UserStatus, error) {

	entity := &orm_model.UserStatus{
		UserID:   status.UserID,
		StatusID: status.StatusID.ToInt(),
	}

	if err := a.db.Create(entity).Error; err != nil {
		return nil, err
	}

	return status, nil
}

// PutUserStatus ユーザーの状態を更新
func (a ActivityRepositoryImpl) PutUserStatus(status *model.UserStatus) (*model.UserStatus, error) {
	userId := status.UserID
	entity := &orm_model.UserStatus{
		StatusID: status.StatusID.ToInt(),
	}
	if err := a.db.Model(&orm_model.UserStatus{}).Where("user_id = ?", userId).Updates(entity).Error; err != nil {
		return nil, err
	}

	return status, nil
}
