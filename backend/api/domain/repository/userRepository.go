package repository

import (
	"work-management-app/domain/model"
)

type UserRepository interface {
	PostUser(user *model.User) (*model.User, error)
	FindUserByUserKey(userKey string) (*model.User, error)
	FindIDByUserKey(userKey string) (id int, err error)
	FindUserStatus(userId int) (*model.UserStatus, error)
	PostUserStatus(status *model.UserStatus) (*model.UserStatus, error)
	PutUserStatus(status *model.UserStatus) (*model.UserStatus, error)
}
