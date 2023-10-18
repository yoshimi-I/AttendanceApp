package usecase

import (
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	"github.com/yoshimi-I/AttendanceApp/usecase/dto/response"
)

type ActivityUsecase interface {
	// レスポンス値にDTOを使う
	AddStudyActivity(activity response.ActivityRequestDTO) (response.ActivityResponseDTO, error)
	UpdateStudyActivity(activity response.ActivityRequestDTO) (response.ActivityResponseDTO, error)
	DeleteActivity(activityID int) error
}

type ActivityUsecaseImpl struct {
	activityRepo repository.ActivtyRepository
}

func (a ActivityUsecaseImpl) AddStudyActivity(activity response.ActivityRequestDTO) (response.ActivityResponseDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (a ActivityUsecaseImpl) UpdateStudyActivity(activity response.ActivityRequestDTO) (response.ActivityResponseDTO, error) {
	//TODO implement me
	panic("implement me")
}

func (a ActivityUsecaseImpl) DeleteActivity(activityID int) error {
	//TODO implement me
	panic("implement me")
}

func NewActivityUsecase(ar repository.ActivtyRepository) ActivityUsecase {
	return &ActivityUsecaseImpl{activityRepo: ar}
}
