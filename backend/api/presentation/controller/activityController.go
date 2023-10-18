package controller

import (
	"github.com/yoshimi-I/AttendanceApp/usecase"
	"net/http"
)

type ActivityController interface {
	AddActivity() http.HandlerFunc
	UpdateActivity() http.HandlerFunc
	DeleteActivity() http.HandlerFunc
}

type ActivityControllerImpl struct {
	ActivityUsecase usecase.ActivityUsecase
}

func (a ActivityControllerImpl) AddActivity() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (a ActivityControllerImpl) UpdateActivity() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func (a ActivityControllerImpl) DeleteActivity() http.HandlerFunc {
	//TODO implement me
	panic("implement me")
}

func NewActivityController(au usecase.ActivityUsecase) ActivityController {
	return &ActivityControllerImpl{ActivityUsecase: au}
}
