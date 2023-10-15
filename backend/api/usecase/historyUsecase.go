package usecase

import (
	"errors"
	"github.com/yoshimi-I/AttendanceApp/domain/model"
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	"time"
)

// まずは扱う関数のinterfaceを実装
type HistoryUsecase interface {
	getStudyActivity(date string) (*model.Activities, error)
	getAllStudyActivity() (*[]model.Activities, error)
}

// 構造体を実装
type HistoryUsecaseImpl struct {
	studyHistoryRepo repository.HistoryRepository
}

// 構造体にinterfaceの関数を実装
func (h HistoryUsecaseImpl) getStudyActivity(date string) (*model.Activities, error) {
	_, err := time.Parse("2006-01-02", date)
	if err != nil {
		return nil, errors.New("invalid date format")
	}
	return &model.Activities{}, nil
}

func (h HistoryUsecaseImpl) getAllStudyActivity() (*[]model.Activities, error) {
	//TODO implement me
	panic("implement me")
}

// 関数を実装した構造体をnewする関数を実装,またこのとき返り値はinterface
func NewHistoryUsecase(hr repository.HistoryRepository) HistoryUsecase {
	return &HistoryUsecaseImpl{studyHistoryRepo: hr}
}
