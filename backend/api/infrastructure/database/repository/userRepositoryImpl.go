package repository

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	"work-management-app/infrastructure/database/orm"
)

type UserRepositoryImpl struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (u UserRepositoryImpl) PostUser(user *model.User, tx repository.Transaction) (*model.User, error) {
	// 受け取ったトランザクションをormに整形
	conn := ConvertOrm(tx)
	if conn != nil {
		conn = u.db
	}

	entity := &orm_model.User{
		Name:    user.Name,
		Email:   user.Email,
		UserKey: user.UserKey,
	}

	if err := conn.Create(entity).Error; err != nil {
		return nil, err
	}

	return &model.User{
		Id:      entity.ID,
		Name:    entity.Name,
		Email:   entity.Email,
		UserKey: entity.UserKey,
	}, nil
}

func (u UserRepositoryImpl) FindUserByUserKey(userKey string) (*model.User, error) {
	var user orm_model.User

	result := u.db.Where("user_key = ?", userKey).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &model.User{
		Id:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		UserKey: user.UserKey,
	}, nil
}

func (u UserRepositoryImpl) FindIDByUserKey(userKey string) (id int, err error) {
	var user orm_model.User
	if result := u.db.Where("user_key = ?", userKey).First(&user); result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, result.Error
		}
		return 0, result.Error
	}
	return user.ID, nil
}

// FindUserStatus 現在のユーザーの状態を確認
func (u UserRepositoryImpl) FindUserStatus(userID int) (*model.UserStatus, error) {
	var status orm_model.UserStatus

	if err := u.db.Where("user_id = ?", userID).First(&status).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no record found with UserID: %d", userID)
		}
		return nil, err
	}

	res := &model.UserStatus{
		UserID:   status.UserID,
		StatusId: model.IntToStatusEnum(status.StatusID),
	}

	return res, nil

}

// PostUserStatus ユーザーの状態を新規登録
func (u UserRepositoryImpl) PostUserStatus(status *model.UserStatus, tx repository.Transaction) (*model.UserStatus, error) {

	// 受け取ったトランザクションをormに整形
	conn := ConvertOrm(tx)
	if conn != nil {
		conn = u.db
	}

	entity := &orm_model.UserStatus{
		UserID:   status.UserID,
		StatusID: status.StatusId.ToInt(),
	}

	if err := conn.Create(entity).Error; err != nil {
		return nil, err
	}

	return status, nil
}

// PutUserStatus ユーザーの状態を更新
func (u UserRepositoryImpl) PutUserStatus(status *model.UserStatus, tx repository.Transaction) (*model.UserStatus, error) {

	// 受け取ったトランザクションをormに整形
	conn := ConvertOrm(tx)
	if conn != nil {
		conn = u.db
	}

	userId := status.UserID
	entity := &orm_model.UserStatus{
		StatusID: status.StatusId.ToInt(),
	}
	if err := conn.Model(&orm_model.UserStatus{}).Where("user_id = ?", userId).Updates(entity).Error; err != nil {
		return nil, err
	}

	return status, nil
}
