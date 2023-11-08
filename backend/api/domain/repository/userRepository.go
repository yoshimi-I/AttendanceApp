package repository

import (
	"work-management-app/domain/model"
)

type UserRepository interface {
	PostUser(user *model.User) (*model.User, error)
	FindUserByUserKey(userKey string) (*model.User, error)
	FindIDByUserKey(userKey string) (id int, err error)
}
