package usecase

import (
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	"net/http"
)

// まずは扱う関数のinterfaceを実装
type StudyHistoryUsecase interface {
	getStudyActivity() http.HandlerFunc
}

type StudyHistoryUsecaseImpl struct {
	studyHistoryRepo repository.StudyHistoryRepository
}

func NewStudyHistoryUsecase(studyHistoryRepo repository.StudyHistoryRepoImpl) StudyHistoryUsecase {
	return &StudyHistoryUsecaseImpl
}
