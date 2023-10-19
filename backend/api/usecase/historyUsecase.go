package usecase

import (
	"fmt"
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	"github.com/yoshimi-I/AttendanceApp/usecase/dto"
)

// まずは扱う関数のinterfaceを実装
type HistoryUsecase interface {
	GetStudyActivityByData(date string) (*dto.HistoryByDateDto, error)
	GetAllStudyHistory() (*dto.ALlHistoryDto, error)
}

// 構造体を実装
type HistoryUsecaseImpl struct {
	hr repository.HistoryRepository
}

func (h HistoryUsecaseImpl) GetAllStudyHistory() (*dto.ALlHistoryDto, error) {
	r, err := h.hr.GetAllHistory()
	if err != nil {
		fmt.Errorf("ミスです")
	}
}

func (h HistoryUsecaseImpl) GetStudyActivityByData(date string) (*dto.HistoryByDateDto, error) {
	r := h.hr.GetHistoryByDate(date)
	return r
}

// 関数を実装した構造体をnewする関数を実装,またこのとき返り値はinterface
func NewHistoryUsecase(hr repository.HistoryRepository) HistoryUsecase {
	return &HistoryUsecaseImpl{hr: hr}
}
