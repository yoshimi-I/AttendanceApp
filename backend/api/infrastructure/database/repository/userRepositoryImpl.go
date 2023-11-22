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

func (u UserRepositoryImpl) PostUser(user *model.User) (*model.User, error) {
	entity := &orm_model.User{
		Name:    user.Name,
		Email:   user.Email,
		UserKey: user.UserKey,
	}

	if err := u.db.Create(entity).Error; err != nil {
		return nil, err
	}

	return &model.User{
		Id:      entity.Id,
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
		Id:      user.Id,
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
	return user.Id, nil
}

// FindUserStatus 現在のユーザーの状態を確認
func (u UserRepositoryImpl) FindUserStatus(userID int) (*model.UserStatus, error) {
	var status orm_model.UserStatus

	if err := u.db.Where("user_id = ?", userID).First(&status).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("no record found with UserId: %d", userID)
		}
		return nil, err
	}

	res := &model.UserStatus{
		UserId:   status.UserId,
		StatusId: model.IntToStatusEnum(status.StatusId),
	}

	return res, nil

}

// PostUserStatus ユーザーの状態を新規登録
func (u UserRepositoryImpl) PostUserStatus(status *model.UserStatus) (*model.UserStatus, error) {

	entity := &orm_model.UserStatus{
		UserId:   status.UserId,
		StatusId: status.StatusId.ToInt(),
	}

	if err := u.db.Create(entity).Error; err != nil {
		return nil, err
	}

	return status, nil
}

// PutUserStatus ユーザーの状態を更新
func (u UserRepositoryImpl) PutUserStatus(status *model.UserStatus) (*model.UserStatus, error) {
	userId := status.UserId
	entity := &orm_model.UserStatus{
		StatusId: status.StatusId.ToInt(),
	}
	if err := u.db.Model(&orm_model.UserStatus{}).Where("user_id = ?", userId).Updates(entity).Error; err != nil {
		return nil, err
	}

	return status, nil
}
