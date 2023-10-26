package usecase

import (
	"fmt"
	"github.com/yoshimi-I/AttendanceApp/domain/model"
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	"github.com/yoshimi-I/AttendanceApp/usecase/dto/request"
	"github.com/yoshimi-I/AttendanceApp/usecase/dto/response"
	"log"
	"time"
)

type ActivityUsecase interface {
	AddStartActivity(activity *request.ActivityStartRequestDTO) (*response.ActivityResponseDTO, error)
	AddEndActivity(activity *request.ActivityEditRequestDTO, id int) (*response.ActivityResponseDTO, error)
	Update(activity *request.ActivityEditRequestDTO, id int) (*response.ActivityResponseDTO, error)
	DeleteByActivityID(activityID int) error
}

type ActivityUsecaseImpl struct {
	ar repository.ActivityRepository
}

func (a ActivityUsecaseImpl) AddStartActivity(activity *request.ActivityStartRequestDTO) (*response.ActivityResponseDTO, error) {

	var res *model.Attendance
	var err error
	var attendance *model.Attendance

	attendance = &model.Attendance{
		UserID:         activity.UserID,
		AttendanceType: activity.AttendanceType,
		StartTime:      activity.StartTime,
		EndTime:        activity.StartTime, // nilではなく開始時刻を使用
		Date:           activity.Date(),
		Year:           activity.Year(),
	}

	res, err = a.ar.PostStartActivity(attendance)
	if err != nil {
		log.Printf("Failed to post start activity: %v", err)
		return nil, fmt.Errorf("failed to post start activity: %w", err)
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		ID:             res.ID,
		UserID:         res.UserID,
		AttendanceType: response.ConvertActivityTime(res.AttendanceType),
		StartTime:      res.StartTime,
		EndTime:        res.EndTime,
		Year:           res.Year,
		Date:           res.Date,
	}
	return responseDTO, nil
}

func (a ActivityUsecaseImpl) AddEndActivity(activity *request.ActivityEditRequestDTO, id int) (*response.ActivityResponseDTO, error) {

	var res *model.Attendance
	var attendance *model.Attendance
	record, err := a.ar.FindActivity(id)

	if err != nil {
		log.Printf("Can't call repository FindActivity: %v", err)
		return nil, err
	}

	// 日付を跨いだ場合の処理を記載
	if record.Date != activity.Date() {
		var firstAttendance *model.Attendance
		dayStartTime := activity.StartTime.Truncate(time.Hour)                               // 0:00
		dayEndTime := activity.StartTime.Truncate(time.Hour).Add(time.Hour*24 - time.Second) // 23:59

		firstAttendance = &model.Attendance{
			ID:      id,
			EndTime: dayEndTime, // 23:59を入れる
		}

		secondAttendance := &model.Attendance{
			UserID:         activity.UserID,
			AttendanceType: activity.AttendanceType,
			StartTime:      dayStartTime, // 0:00
			EndTime:        activity.EndTime,
			Date:           activity.Date(),
			Year:           activity.Year(),
		}

		// 2つの値をDBに保存
		_, err = a.ar.PostEndActivity(firstAttendance)
		res, err = a.ar.PostStartActivity(secondAttendance)
		if err != nil {
			log.Printf("Failed to crossing over to the next day: %v", err)
			return nil, fmt.Errorf("failed to crossing over to the next day: %w", err)
		}

	} else {
		attendance = &model.Attendance{
			ID:      id,
			EndTime: activity.EndTime,
		}
		res, err = a.ar.PostEndActivity(attendance)
		if err != nil {
			log.Printf("Failed to post end activity: %v", err)
			return nil, fmt.Errorf("failed to post end activity: %w", err)
		}
	}
	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		ID:             res.ID,
		UserID:         res.UserID,
		AttendanceType: response.ConvertActivityTime(res.AttendanceType),
		StartTime:      res.StartTime,
		EndTime:        res.EndTime,
		Year:           res.Year,
		Date:           res.Date,
	}
	return responseDTO, nil
}

func (a ActivityUsecaseImpl) Update(activity *request.ActivityEditRequestDTO, id int) (*response.ActivityResponseDTO, error) {
	_, err := a.ar.FindActivity(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find existing activity: %v", err)
	}
	attendance := &model.Attendance{
		ID:        id,
		StartTime: activity.StartTime,
		EndTime:   activity.EndTime,
	}
	res, err := a.ar.PostEndActivity(attendance)
	if err != nil {
		return nil, fmt.Errorf("failed to update activity end time: %v", err)
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		ID:             res.ID,
		UserID:         res.UserID,
		AttendanceType: response.ConvertActivityTime(res.AttendanceType),
		StartTime:      res.StartTime,
		EndTime:        res.EndTime,
		Year:           res.Year,
		Date:           res.Date,
	}
	return responseDTO, nil
}

func (a ActivityUsecaseImpl) DeleteByActivityID(activityID int) error {
	_, err := a.ar.FindActivity(activityID)
	if err != nil {
		return fmt.Errorf("failed to find existing activity: %v", err)
	}
	err = a.ar.DeleteActivity(activityID)
	if err != nil {
		return fmt.Errorf("failed to delete activity: %v", err)
	}

	return nil

}

func NewActivityUsecase(ar repository.ActivityRepository) ActivityUsecase {
	return &ActivityUsecaseImpl{ar: ar}
}
