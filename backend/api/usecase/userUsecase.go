package usecase

import (
	"log"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	"work-management-app/usecase/dto/request"
	"work-management-app/usecase/dto/response"
)

type UserUsecase interface {
	AddUser(user *request.UserDTO) (*response.UserDTO, error)
	UserByUserKey(userKey string) (*response.UserDTO, error)
	IDByUserKey(userKey string) (id int, err error)
}

type UserUsecaseImpl struct {
	ur repository.UserRepository
}

func NewUserUsecase(ur repository.UserRepository) UserUsecase {
	return &UserUsecaseImpl{ur: ur}
}
func (u UserUsecaseImpl) AddUser(user *request.UserDTO) (*response.UserDTO, error) {
	addUser := &model.User{
		Name:    user.Name,
		Email:   user.Email,
		UserKey: user.UserKey,
	}
	res, err := u.ur.PostUser(addUser)
	if err != nil {
		log.Printf("Failed to adduser in usecase: %v", err)
		return nil, err
	}

	// DTOに詰め替え
	responseDTO := &response.UserDTO{
		Id:         res.ID,
		Name:       res.Name,
		Email:      res.Email,
		UserKeyKey: res.UserKey,
	}
	return responseDTO, nil

}

func (u UserUsecaseImpl) UserByUserKey(userKey string) (*response.UserDTO, error) {
	res, err := u.ur.FindUserByUserKey(userKey)
	if err != nil {
		log.Printf("Faild to UserById in usecase : %v", err)
		return nil, err
	}

	// DTO詰め替え
	responseDTO := &response.UserDTO{
		Id:         res.ID,
		Name:       res.Name,
		Email:      res.Email,
		UserKeyKey: res.UserKey,
	}
	return responseDTO, nil
}

func (u UserUsecaseImpl) IDByUserKey(userKey string) (id int, err error) {
	res, err := u.ur.FindIDByUserKey(userKey)
	if err != nil {
		log.Printf("Faild to IDByUserKey in usecase : %v", err)
		return 0, err
	}
	return res, nil
}
