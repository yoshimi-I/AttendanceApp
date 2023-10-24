package usecase

import (
	"fmt"
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	"github.com/yoshimi-I/AttendanceApp/usecase/dto/response"
)

// まずは扱う関数のinterfaceを実装
type HistoryUsecase interface {
	GetStudyActivityByData(date string) (*response.HistoryByDateDto, error)
	GetAllStudyHistory() (*response.ALlHistoryDto, error)
}

// 構造体を実装
type HistoryUsecaseImpl struct {
	hr repository.HistoryRepository
}

func (h HistoryUsecaseImpl) GetAllStudyHistory() (*response.ALlHistoryDto, error) {
	r, _ := h.hr.GetAllHistory(1)
	if r != nil {
		fmt.Println("やあ")
	}
	//if err != nil {
	//	fmt.Errorf("ミスです")
	//}
	return nil, nil
}

func (h HistoryUsecaseImpl) GetStudyActivityByData(date string) (*response.HistoryByDateDto, error) {
	//r := h.hr.GetHistoryByDate(date)
	return nil, nil
}

// 関数を実装した構造体をnewする関数を実装,またこのとき返り値はinterface
func NewHistoryUsecase(hr repository.HistoryRepository) HistoryUsecase {
	return &HistoryUsecaseImpl{hr: hr}
}
