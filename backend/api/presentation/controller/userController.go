package controller

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"work-management-app/usecase"
	"work-management-app/usecase/dto/request"
)

type UserController interface {
	CreateUser() http.HandlerFunc
	GetStatus() http.HandlerFunc
}

type UserControllerImpl struct {
	uu usecase.UserUsecase
}

func NewUserController(uu usecase.UserUsecase) UserController {
	return &UserControllerImpl{uu: uu}
}
func (u UserControllerImpl) CreateUser() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// bodyから受け取る型
		var user request.UserDTO

		// bodyを取得
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			log.Println("Can't get body")
			http.Error(w, "Invalid request body", http.StatusBadRequest)
			return
		}

		res, err := u.uu.AddUser(&user)
		if err != nil {
			log.Println("Error in AddUserUsecase")
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			log.Printf("Can't encode json: %v", err)
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}

}

func (u UserControllerImpl) GetStatus() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// パラメータ取得
		userKey := chi.URLParam(r, "userKey")
		res, err := u.uu.UserStatusByUserKey(userKey)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
		}

		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(res); err != nil {
			http.Error(w, "Failed to encode response", http.StatusInternalServerError)
			return
		}
	}

}
