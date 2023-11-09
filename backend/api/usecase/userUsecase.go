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
	ar repository.ActivityRepository
}

func NewUserUsecase(ur repository.UserRepository, ar repository.ActivityRepository) UserUsecase {
	return &UserUsecaseImpl{ur: ur, ar: ar}
}

func (u UserUsecaseImpl) AddUser(user *request.UserDTO) (*response.UserDTO, error) {

	var res *model.User
	userKey := user.UserKey
	// 重複を確認
	findUser, err := u.ur.FindUserByUserKey(userKey)

	if err != nil {
		log.Printf("Failed to adduser in repository: %v", err)
		return nil, err
	}

	// userがまだ登録されていない場合のみDBに保存
	if findUser != nil {
		res = findUser

	} else {
		addUser := &model.User{
			Name:    user.Name,
			Email:   user.Email,
			UserKey: userKey,
		}

		// DBに保存
		res, err = u.ur.PostUser(addUser)
		if err != nil {
			return nil, err
		}

		// その後ユーザーの状態を保持(最初はFinish)
		addUserStatus := &model.UserStatus{
			UserID:   res.ID,
			StatusID: model.Finish,
		}
		log.Printf("Setting initial user status for user: %s", userKey)
		_, err := u.ar.PostUserStatus(addUserStatus)
		if err != nil {
			return nil, err
		}

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
