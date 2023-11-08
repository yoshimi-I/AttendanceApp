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

	return user, nil
}

func (u UserRepositoryImpl) FindByID(id int) (*model.User, error) {
	var user orm_model.User

	if err := u.db.First(&user, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, err
		}
		return nil, err
	}

	res := &model.User{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		UserKey: user.UserKey,
	}

	return res, nil
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
