package repository

import (
	"errors"
	"gorm.io/gorm"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	orm_model "work-management-app/infrastructure/orm"
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
		ID:      entity.ID,
		Name:    entity.Name,
		Email:   entity.Email,
		UserKey: entity.UserKey,
	}, nil
}

func (u *UserRepositoryImpl) FindUserByUserKey(userKey string) (*model.User, error) {
	var user orm_model.User

	result := u.db.Where("user_key = ?", userKey).First(&user)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil, nil
	}

	if result.Error != nil {
		return nil, result.Error
	}

	return &model.User{
		ID:      user.ID,
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
