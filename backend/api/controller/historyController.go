package controller

import (
	"github.com/yoshimi-I/AttendanceApp/usecase"
	"net/http"
)

type HistoryController interface {
	GetAllHistory() http.HandlerFunc
	GetHistoryByDate() http.HandlerFunc
}

type HistoryControllerImpl struct {
	// ここで使うusecaseを全て実装
	historyUsecasse usecase.HistoryUsecase
}

func (h HistoryControllerImpl) GetAllHistory() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		return
	}
}

func (h HistoryControllerImpl) GetHistoryByDate() http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		return
	}
}

func NewHisoryController(hu usecase.HistoryUsecase) HistoryController {
	return &HistoryControllerImpl{historyUsecasse: hu}
}
