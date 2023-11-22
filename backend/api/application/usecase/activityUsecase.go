package usecase

import (
	"fmt"
	"log"
	"time"
	"work-management-app/application/dto/request"
	"work-management-app/application/dto/response"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	"work-management-app/utility"
)

type ActivityUsecase interface {
	AddStarWork(work *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error)
	AddEndWork(work *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error)
	AddStartBreak(breakInfo *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error)
	AddEndBreak(breakInfo *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error)
	Update(activity *request.ActivityEditRequestDTO) (*response.ActivityResponseDTO, error)
	DeleteByActivityID(activity *request.ActivityDeleteRequestDTO) error
}

type ActivityUsecaseImpl struct {
	ar repository.ActivityRepository
	ur repository.UserRepository
	hr repository.HistoryRepository
}

func NewActivityUsecase(ar repository.ActivityRepository, ur repository.UserRepository, hr repository.HistoryRepository) ActivityUsecase {
	return &ActivityUsecaseImpl{
		ar: ar,
		ur: ur,
		hr: hr,
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
	nowUserStatus, err := a.ur.FindUserStatus(userID)
	if err != nil {
		log.Println("usr_id not found")
		return nil, err
	}

	// 終了の状態の時のみ,作業を開始できる
	if nowUserStatus.StatusId != model.Finish {
		log.Println(nowUserStatus.StatusId.ToString())
		return nil, fmt.Errorf("作業の開始は現在行えません")
	}

	if err != nil {
		return nil, err
	}

	// ユーザーの状態(作業中)を登録
	updateUserStatus := &model.UserStatus{
		UserId:   userID,
		StatusId: model.Work,
	}
	userStatus, err := a.ur.PutUserStatus(updateUserStatus)
	if err != nil {
		return nil, err
	}

	// 作業の登録
	attendance := &model.Attendance{
		UserId:         userID,
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
		Id:             res.Id,
		AttendanceType: "work_start",
		Time:           res.Time,
		Year:           res.Year,
		Date:           res.Date,
		Status:         userStatus.StatusId.ToString(),
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
	nowUserStatus, err := a.ur.FindUserStatus(userID)
	if err != nil {
		log.Println("user_status is not found")
		return nil, err
	}

	if nowUserStatus.StatusId != model.Work {
		return nil, fmt.Errorf("作業の終了は現在行えません")
	}

	if err != nil {
		return nil, err
	}

	// ユーザーの状態(終了)を登録
	updateUserStatus := &model.UserStatus{
		UserId:   userID,
		StatusId: model.Finish,
	}

	userStatus, err := a.ur.PutUserStatus(updateUserStatus)
	if err != nil {
		return nil, err
	}

	// 作業の登録
	attendance := &model.Attendance{
		UserId:         userID,
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
		Id:             res.Id,
		AttendanceType: "work_end",
		Time:           res.Time,
		Year:           res.Year,
		Date:           res.Date,
		Status:         userStatus.StatusId.ToString(),
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
	nowUserStatus, err := a.ur.FindUserStatus(userID)
	if err != nil {
		log.Println("user_status is not found")

		return nil, err
	}

	if nowUserStatus.StatusId != model.Work {
		return nil, fmt.Errorf("休憩の開始は現在行えません")
	}

	if err != nil {
		return nil, err
	}

	// ユーザーの状態(休憩中)を登録
	updateUserStatus := &model.UserStatus{
		UserId:   userID,
		StatusId: model.Break,
	}
	userStatus, err := a.ur.PutUserStatus(updateUserStatus)
	if err != nil {
		return nil, err
	}

	// 作業の登録
	attendance := &model.Attendance{
		UserId:         userID,
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
		Id:             res.Id,
		AttendanceType: "break_start",
		Time:           res.Time,
		Year:           res.Year,
		Date:           res.Date,
		Status:         userStatus.StatusId.ToString(),
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
	nowUserStatus, err := a.ur.FindUserStatus(userID)
	if err != nil {
		log.Println("user_status is not found")
		return nil, err
	}

	if nowUserStatus.StatusId != model.Break {
		return nil, fmt.Errorf("休憩の終了は現在行えません")
	}

	if err != nil {
		return nil, err
	}

	// ユーザーの状態(作業中)を登録
	updateUserStatus := &model.UserStatus{
		UserId:   userID,
		StatusId: model.Work,
	}
	userStatus, err := a.ur.PutUserStatus(updateUserStatus)
	if err != nil {
		return nil, err
	}

	// 作業の登録
	attendance := &model.Attendance{
		UserId:         userID,
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
		Id:             res.Id,
		AttendanceType: "break_start",
		Time:           res.Time,
		Year:           res.Year,
		Date:           res.Date,
		Status:         userStatus.StatusId.ToString(),
	}
	return responseDTO, nil
}

// Update 作業,休憩の修正
func (a ActivityUsecaseImpl) Update(activity *request.ActivityEditRequestDTO) (*response.ActivityResponseDTO, error) {
	activityID := activity.ActivityId
	userKey := activity.UserKey
	newTime := activity.Time

	//userKeyからuserIDを取得
	userID, err := a.ur.FindIDByUserKey(userKey)

	// activityが存在するかどうかを確認
	record, err := a.ar.FindActivity(activityID)
	if err != nil {
		return nil, utility.NotFoundError{Message: "failed to find existing activity"}
	}

	// 編集処理をする人が本当に本人かどうかを確認
	if userID != record.UserId {
		return nil, utility.AuthenticationError{Message: "user authentication failed"}
	}

	// Dateからユーザーの行動を取得
	dateStr := record.Date
	historyByDate, err := a.hr.ReadHistoryByDate(userID, dateStr)

	// 編集時データのバリデーションチェック
	var beforeTime, afterTime time.Time
	for i, history := range historyByDate {
		if history.Id == activityID {
			if i == 0 {
				// 最初の値かどうかを確認
				// 開始時刻を代入
				beforeTime = utility.StartTime()
				afterTime = historyByDate[i+1].Time
			} else if i == len(historyByDate) {
				// 最後の値かどうか
				// 現在の日本時間を代入
				beforeTime = historyByDate[i-1].Time
				afterTime = utility.NowTime()
			} else {
				beforeTime = historyByDate[i-1].Time
				afterTime = historyByDate[i+1].Time
			}
		}
	}

	if !utility.IsTimeInRange(beforeTime, newTime, afterTime) {
		return nil, utility.BadRequestError{}
	}

	// 編集するデータを詰め替える
	attendance := &model.Attendance{
		Id:   activityID,
		Time: newTime,
	}

	res, err := a.ar.PutActivity(attendance)
	if err != nil {
		return nil, fmt.Errorf("failed to update activity end time: %v", err)
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		Id:             res.Id,
		AttendanceType: response.ConvertActivityTime(record.AttendanceType),
		Time:           res.Time,
		Year:           record.Year,
		Date:           record.Date,
	}
	return responseDTO, nil
}

// DeleteByActivityID 作業,休憩の削除
func (a ActivityUsecaseImpl) DeleteByActivityID(activity *request.ActivityDeleteRequestDTO) error {
	activityID := activity.ActivityId
	userKey := activity.UserKey
	var err error

	//userKeyからuserIDを取得
	userID, err := a.ur.FindIDByUserKey(userKey)

	// activityが存在するかどうかを確認
	record, err := a.ar.FindActivity(activityID)
	if err != nil {
		return utility.NotFoundError{Message: "failed to find existing activity"}
	}

	// 編集処理をする人の権限確認
	if userID != record.UserId {
		return utility.AuthenticationError{Message: "user authentication failed"}
	}

	// 削除時データのバリデーションチェック
	// 削除はその日の一番新しいものしかできないようにする(整合性を保つため)
	dateStr := record.Date // "2023-12-25"
	historyByDate, err := a.hr.ReadHistoryByDate(userID, dateStr)

	// 当日の削除の場合は現在の状態も更新
	// 別の日の削除の場合は状態の変更はしない
	for i, history := range historyByDate {

		// 削除したいデータが日付ごとの最新のデータでない場合はエラーを返す
		if activityID == history.Id {
			if i != len(historyByDate)-1 {
				return utility.ForbiddenError{}
			}
			if dateStr == history.Date {

				// 更新したい処理を確認
				var newAction model.StatusEnum
				// 削除したい処理を確認
				deleteAction := history.AttendanceType
				switch deleteAction {
				case model.WorkStart:
					newAction = model.Finish
				case model.WorkEnd:
					newAction = model.Work
				case model.BreakStart:
					newAction = model.Work
				case model.BreakEnd:
					newAction = model.Break
				default:
					return err
				}
				// 現在の状態の更新を行う
				updateUserStatus := &model.UserStatus{
					UserId:   userID,
					StatusId: newAction,
				}
				_, err := a.ur.PutUserStatus(updateUserStatus)
				if err != nil {
					return err
				}
				// 削除処理を行う
				err = a.ar.DeleteActivity(activityID)
			} else {
				// 削除処理を行う
				err = a.ar.DeleteActivity(activityID)
			}
		}

	}
	return nil

}
