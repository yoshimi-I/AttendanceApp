package usecase

import (
	"fmt"
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	"github.com/yoshimi-I/AttendanceApp/usecase/dto/response"
)

// まずは扱う関数のinterfaceを実装
type HistoryUsecase interface {
	GetStudyHistory(date string) (*response.HistoryByDateDto, error)
	GetAllStudyHistory() (*response.ALlHistoryDto, error)
}

// 構造体を実装
type HistoryUsecaseImpl struct {
	hr repository.HistoryRepository
}

func (h HistoryUsecaseImpl) GetAllStudyHistory() (*[]interface{}, error) {
	//TODO implement me
	panic("implement me")
}

// 構造体にinterfaceの関数を実装
func (h HistoryUsecaseImpl) GetStudyHistory(date string) (*response.HistoryByDateDto, error) {
	response := &dto.{
		Type: "勉強開始",

		Timestamp: parsedDate,
	}

	return response, nil
}

func (h HistoryUsecaseImpl) GetAllStudyAHistory() (*response.ALlHistoryDto, error) {
	activities, err := h.hr.GetAllHistory()
	if err != nil {
		return nil, fmt.Errorf("failed to get all study activities: %w", err)
	}
	return &activities, err
}

// 関数を実装した構造体をnewする関数を実装,またこのとき返り値はinterface
func NewHistoryUsecase(hr repository.HistoryRepository) HistoryUsecase {
	return &HistoryUsecaseImpl{hr: hr}
}
