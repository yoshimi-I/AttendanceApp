package usecase

import (
	"fmt"
	"log"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	"work-management-app/usecase/dto/request"
	"work-management-app/usecase/dto/response"
)

type ActivityUsecase interface {
	AddStarWork(work *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error)
	AddEndWork(work *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error)
	AddStartBreak(breakInfo *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error)
	AddEndBreak(breakInfo *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error)
	Update(activity *request.ActivityEditRequestDTO, id int) (*response.ActivityResponseDTO, error)
	DeleteByActivityID(activityID int) error
}

type ActivityUsecaseImpl struct {
	ar repository.ActivityRepository
	ur repository.UserRepository
}

func NewActivityUsecase(ar repository.ActivityRepository, ur repository.UserRepository) ActivityUsecase {
	return &ActivityUsecaseImpl{
		ar: ar,
		ur: ur,
	}
}

// AddStarWork　作業の開始を登録
func (a ActivityUsecaseImpl) AddStarWork(work *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error) {

	var res *model.Attendance
	var err error
	userKey := work.UserKey

	// userKeyからuserIdを指定
	userID, err := a.ur.FindIDByUserKey(userKey)
	if err != nil {
		log.Println("usr_id not found")
		return nil, err
	}

	// 現在の状態をまずは取得
	nowUserStatus, err := a.ar.FindUserStatus(userID)
	if err != nil {
		log.Println("usr_id not found")
		return nil, err
	}
	// 終了の状態の時のみ,作業を開始できる
	if nowUserStatus.StatusID != model.Finish {
		log.Println(nowUserStatus.StatusID.ToString())
		return nil, fmt.Errorf("作業の開始は現在行えません")
	}

	if err != nil {
		return nil, err
	}

	// ユーザーの状態(作業中)を登録
	updateUserStatus := &model.UserStatus{
		UserID:   userID,
		StatusID: model.Work,
	}
	userStatus, err := a.ar.PutUserStatus(updateUserStatus)
	if err != nil {
		return nil, err
	}

	// 作業の登録
	attendance := &model.Attendance{
		UserID:         userID,
		AttendanceType: model.WorkStart,
		Time:           work.Time,
		Date:           work.Date(),
		Year:           work.Year(),
	}

	res, err = a.ar.PostActivity(attendance)
	if err != nil {
		log.Printf("Failed to post start activity: %v", err)
		return nil, fmt.Errorf("failed to post start activity: %w", err)
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		ID:             res.ID,
		AttendanceType: "work_start",
		Time:           res.Time,
		Year:           res.Year,
		Date:           res.Date,
		Status:         userStatus.StatusID.ToString(),
	}
	return responseDTO, nil
}

// AddEndWork 作業の終了を登録
func (a ActivityUsecaseImpl) AddEndWork(work *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error) {

	var res *model.Attendance
	userKey := work.UserKey

	// userKeyからuserIdを指定
	userID, err := a.ur.FindIDByUserKey(userKey)
	if err != nil {
		log.Println("usr_id not found")
		return nil, err
	}

	// 現在の状態をまずは取得
	nowUserStatus, err := a.ar.FindUserStatus(userID)
	if err != nil {
		log.Println("user_status is not found")
		return nil, err
	}

	if nowUserStatus.StatusID != model.Work {
		return nil, fmt.Errorf("作業の終了は現在行えません")
	}

	if err != nil {
		return nil, err
	}

	// ユーザーの状態(終了)を登録
	updateUserStatus := &model.UserStatus{
		UserID:   userID,
		StatusID: model.Finish,
	}

	userStatus, err := a.ar.PutUserStatus(updateUserStatus)
	if err != nil {
		return nil, err
	}

	// 作業の登録
	attendance := &model.Attendance{
		UserID:         userID,
		AttendanceType: model.WorkEnd,
		Time:           work.Time,
		Date:           work.Date(),
		Year:           work.Year(),
	}

	res, err = a.ar.PostActivity(attendance)
	if err != nil {
		log.Printf("Failed to post start activity: %v", err)
		return nil, fmt.Errorf("failed to post start activity: %w", err)
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		ID:             res.ID,
		AttendanceType: "work_end",
		Time:           res.Time,
		Year:           res.Year,
		Date:           res.Date,
		Status:         userStatus.StatusID.ToString(),
	}
	return responseDTO, nil
}

// AddStartBreak　休憩の開始を登録
func (a ActivityUsecaseImpl) AddStartBreak(breakInfo *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error) {

	var res *model.Attendance
	var err error
	userKey := breakInfo.UserKey

	// userKeyからuserIdを指定
	userID, err := a.ur.FindIDByUserKey(userKey)
	if err != nil {
		log.Println("usr_id not found")
		return nil, err
	}

	// 現在の状態をまずは取得
	nowUserStatus, err := a.ar.FindUserStatus(userID)
	if err != nil {
		log.Println("user_status is not found")

		return nil, err
	}

	if nowUserStatus.StatusID != model.Work {
		return nil, fmt.Errorf("休憩の開始は現在行えません")
	}

	if err != nil {
		return nil, err
	}

	// ユーザーの状態(休憩中)を登録
	updateUserStatus := &model.UserStatus{
		UserID:   userID,
		StatusID: model.Break,
	}
	userStatus, err := a.ar.PutUserStatus(updateUserStatus)
	if err != nil {
		return nil, err
	}

	// 作業の登録
	attendance := &model.Attendance{
		UserID:         userID,
		AttendanceType: model.BreakStart,
		Time:           breakInfo.Time,
		Date:           breakInfo.Date(),
		Year:           breakInfo.Year(),
	}

	// 休憩時間を登録
	res, err = a.ar.PostActivity(attendance)
	if err != nil {
		log.Printf("Failed to post start break: %v", err)
		return nil, fmt.Errorf("failed to post start break: %w", err)
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		ID:             res.ID,
		AttendanceType: "break_start",
		Time:           res.Time,
		Year:           res.Year,
		Date:           res.Date,
		Status:         userStatus.StatusID.ToString(),
	}
	return responseDTO, nil
}

// AddEndBreak 休憩の終了を登録
func (a ActivityUsecaseImpl) AddEndBreak(breakInfo *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error) {

	var res *model.Attendance
	userKey := breakInfo.UserKey

	// userKeyからuserIdを指定
	userID, err := a.ur.FindIDByUserKey(userKey)
	if err != nil {
		log.Println("usr_id not found")
		return nil, err
	}

	// 現在の状態をまずは取得
	nowUserStatus, err := a.ar.FindUserStatus(userID)
	if err != nil {
		log.Println("user_status is not found")
		return nil, err
	}

	if nowUserStatus.StatusID != model.Break {
		return nil, fmt.Errorf("休憩の終了は現在行えません")
	}

	if err != nil {
		return nil, err
	}

	// ユーザーの状態(作業中)を登録
	updateUserStatus := &model.UserStatus{
		UserID:   userID,
		StatusID: model.Work,
	}
	userStatus, err := a.ar.PutUserStatus(updateUserStatus)
	if err != nil {
		return nil, err
	}

	// 作業の登録
	attendance := &model.Attendance{
		UserID:         userID,
		AttendanceType: model.BreakEnd,
		Time:           breakInfo.Time,
		Date:           breakInfo.Date(),
		Year:           breakInfo.Year(),
	}

	// 休憩時間を登録
	res, err = a.ar.PostActivity(attendance)
	if err != nil {
		log.Printf("Failed to post start break: %v", err)
		return nil, fmt.Errorf("failed to post start break: %w", err)
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		ID:             res.ID,
		AttendanceType: "break_start",
		Time:           res.Time,
		Year:           res.Year,
		Date:           res.Date,
		Status:         userStatus.StatusID.ToString(),
	}
	return responseDTO, nil
}

// Update 作業,休憩の修正
func (a ActivityUsecaseImpl) Update(activity *request.ActivityEditRequestDTO, id int) (*response.ActivityResponseDTO, error) {
	record, err := a.ar.FindActivity(id)
	if err != nil {
		return nil, fmt.Errorf("failed to find existing activity: %v", err)
	}
	attendance := &model.Attendance{
		ID:   id,
		Time: activity.Time,
	}
	res, err := a.ar.PostActivity(attendance)
	if err != nil {
		return nil, fmt.Errorf("failed to update activity end time: %v", err)
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		ID:             res.ID,
		AttendanceType: response.ConvertActivityTime(record.AttendanceType),
		Time:           res.Time,
		Year:           record.Year,
		Date:           record.Date,
	}
	return responseDTO, nil
}

// DeleteByActivityID 作業,休憩の削除
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
