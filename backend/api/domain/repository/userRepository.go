package repository

import (
	"work-management-app/domain/model"
)

type UserRepository interface {
	PostUser(user *model.User) (*model.User, error)
	FindByID(id int) (*model.User, error)
	FindIDByUserKey(userKey string) (id int, err error)
}
