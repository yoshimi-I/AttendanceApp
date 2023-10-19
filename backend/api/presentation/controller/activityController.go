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
	return nil
}

func (a ActivityControllerImpl) UpdateActivity() http.HandlerFunc {
	return nil
}

func (a ActivityControllerImpl) DeleteActivity() http.HandlerFunc {
	return nil
}

func NewActivityController(au usecase.ActivityUsecase) ActivityController {
	return &ActivityControllerImpl{ActivityUsecase: au}
}
