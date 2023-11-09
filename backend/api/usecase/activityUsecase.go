package usecase

import (
	"fmt"
	"log"
	"time"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	"work-management-app/usecase/dto/request"
	"work-management-app/usecase/dto/response"
)

type ActivityUsecase interface {
	AddStarWork(work *request.ActivityStartRequestDTO) (*response.ActivityResponseDTO, error)
	AddEndWork(work *request.ActivityEndRequestDTO, id int) (*response.ActivityResponseDTO, error)
	AddStartBreak(Break *request.ActivityStartRequestDTO) (*response.ActivityResponseDTO, error)
	AddEndBreak(Break *request.ActivityEndRequestDTO, id int) (*response.ActivityResponseDTO, error)
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
func (a ActivityUsecaseImpl) AddStarWork(work *request.ActivityStartRequestDTO) (*response.ActivityResponseDTO, error) {

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
	if nowUserStatus.StatusID.ToString() != "Finish" {
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
		AttendanceType: model.WorkStartEnd,
		StartTime:      work.StartTime,
		EndTime:        work.StartTime, // nilではなく開始時刻を使用
		Date:           work.Date(),
		Year:           work.Year(),
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
		AttendanceType: "work_start",
		StartTime:      res.StartTime,
		EndTime:        res.EndTime,
		Year:           res.Year,
		Date:           res.Date,
		Status:         userStatus.StatusID.ToString(),
	}
	return responseDTO, nil
}

// AddEndWork 作業の終了を登録
func (a ActivityUsecaseImpl) AddEndWork(work *request.ActivityEndRequestDTO, id int) (*response.ActivityResponseDTO, error) {

	var res *model.Attendance
	var attendance *model.Attendance
	userKey := work.UserKey

	record, err := a.ar.FindActivity(id)

	if err != nil {
		log.Printf("Can't call repository FindActivity: %v", err)
		return nil, err
	}

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

	if nowUserStatus.StatusID.ToString() != "Work" {
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

	// 日付を跨いだ場合の処理を記載
	if record.Date != work.Date() {
		var firstAttendance *model.Attendance
		currentDate := time.Now()

		// 0:00を表現
		dayStartTime := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.UTC)

		// 23:59を表現
		dayEndTime := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 23, 59, 0, 0, time.UTC)

		firstAttendance = &model.Attendance{
			ID:        id,
			StartTime: res.StartTime,
			EndTime:   dayEndTime, // 23:59を入れる
		}

		if err != nil {
			return nil, err
		}

		secondAttendance := &model.Attendance{
			AttendanceType: model.WorkStartEnd,
			StartTime:      dayStartTime, // 0:00
			EndTime:        work.EndTime,
			Date:           work.Date(),
			Year:           work.Year(),
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
			ID:        id,
			StartTime: record.StartTime,
			EndTime:   work.EndTime,
		}
		res, err = a.ar.PostEndActivity(attendance)
		if err != nil {
			log.Printf("Failed to post end activity: %v", err)
			return nil, fmt.Errorf("failed to post end activity: %w", err)
		}
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		ID:             record.ID,
		UserID:         record.UserID,
		AttendanceType: "work_end",
		StartTime:      record.StartTime,
		EndTime:        res.EndTime,
		Year:           record.Year,
		Date:           record.Date,
		Status:         userStatus.StatusID.ToString(),
	}
	return responseDTO, nil
}

// AddStartBreak　休憩の開始を登録
func (a ActivityUsecaseImpl) AddStartBreak(Break *request.ActivityStartRequestDTO) (*response.ActivityResponseDTO, error) {

	var res *model.Attendance
	var err error
	userKey := Break.UserKey

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

	if nowUserStatus.StatusID.ToString() != "Work" {
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

	attendance := &model.Attendance{
		UserID:         userID,
		AttendanceType: model.BreakStartEnd,
		StartTime:      Break.StartTime,
		EndTime:        Break.StartTime, // nilではなく開始時刻を使用
		Date:           Break.Date(),
		Year:           Break.Year(),
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
		AttendanceType: "break_start",
		StartTime:      res.StartTime,
		EndTime:        res.EndTime,
		Year:           res.Year,
		Date:           res.Date,
		Status:         userStatus.StatusID.ToString(),
	}
	return responseDTO, nil
}

// AddEndBreak 作業の終了,勤務の終了を登録
func (a ActivityUsecaseImpl) AddEndBreak(Break *request.ActivityEndRequestDTO, id int) (*response.ActivityResponseDTO, error) {

	var res *model.Attendance
	var attendance *model.Attendance
	userKey := Break.UserKey

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

	if nowUserStatus.StatusID.ToString() != "Break" {
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

	record, err := a.ar.FindActivity(id)

	if err != nil {
		log.Printf("Can't call repository FindActivity: %v", err)
		return nil, err
	}

	// 日付を跨いだ場合の処理を記載
	if record.Date != Break.Date() {
		var firstAttendance *model.Attendance
		currentDate := time.Now()

		// 0:00を表現
		dayStartTime := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 0, 0, 0, 0, time.UTC)

		// 23:59を表現
		dayEndTime := time.Date(currentDate.Year(), currentDate.Month(), currentDate.Day(), 23, 59, 0, 0, time.UTC)

		firstAttendance = &model.Attendance{
			ID:        id,
			StartTime: record.StartTime,
			EndTime:   dayEndTime, // 23:59を入れる
		}

		if err != nil {
			return nil, err
		}

		secondAttendance := &model.Attendance{
			AttendanceType: model.BreakStartEnd,
			StartTime:      dayStartTime, // 0:00
			EndTime:        Break.EndTime,
			Date:           Break.Date(),
			Year:           Break.Year(),
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
			ID:        id,
			StartTime: record.StartTime,
			EndTime:   Break.EndTime,
		}
		res, err = a.ar.PostEndActivity(attendance)
		if err != nil {
			log.Printf("Failed to post end activity: %v", err)
			return nil, fmt.Errorf("failed to post end activity: %w", err)
		}
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		ID:             record.ID,
		UserID:         record.UserID,
		AttendanceType: "break_end",
		StartTime:      record.StartTime,
		EndTime:        res.EndTime,
		Year:           record.Year,
		Date:           record.Date,
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
		UserID:         record.UserID,
		AttendanceType: response.ConvertActivityTime(record.AttendanceType),
		StartTime:      res.StartTime,
		EndTime:        res.EndTime,
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
