package service

import (
	"fmt"
	"log"
	"time"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	"work-management-app/utility"
)

type ActivityDomainService interface {
	AddStarWorkTime(userID int, tx repository.Transaction) (*model.Attendance, error)
	AddEndWorkTime(userID int, tx repository.Transaction) (*model.Attendance, error)
	AddStartBreakTime(userID int, tx repository.Transaction) (*model.Attendance, error)
	AddEndBreakTime(userID int, tx repository.Transaction) (*model.Attendance, error)
	EditTime(attendance *model.Attendance, newTime time.Time, tx repository.Transaction) (*model.Attendance, error)
	Delete(attendance *model.Attendance, tx repository.Transaction) (*model.UserStatus, error)
}

type ActivityServiceImpl struct {
	ar repository.ActivityRepository
	hr repository.HistoryRepository
}

func NewActivityDomainService(ar repository.ActivityRepository, hr repository.HistoryRepository) ActivityDomainService {
	return &ActivityServiceImpl{
		ar: ar,
		hr: hr,
	}
}

// AddStarWorkTime　作業の開始を登録
func (a ActivityServiceImpl) AddStarWorkTime(userID int, tx repository.Transaction) (*model.Attendance, error) {

	// 時間の登録
	attendance := &model.Attendance{
		UserID:         userID,
		AttendanceType: model.WorkStart,
		Time:           utility.NowTime(),
		Date:           utility.NowDateStr(),
		Year:           utility.NowYear(),
	}

	res, err := a.ar.PostActivity(attendance, tx)
	if err != nil {
		log.Printf("Failed to post start activity: %v", err)
		return nil, fmt.Errorf("failed to post start activity: %w", err)
	}

	return res, nil
}

// AddEndWorkTime  作業の終了を登録
func (a ActivityServiceImpl) AddEndWorkTime(userID int, tx repository.Transaction) (*model.Attendance, error) {

	// 時間の登録
	attendance := &model.Attendance{
		UserID:         userID,
		AttendanceType: model.WorkEnd,
		Time:           utility.NowTime(),
		Date:           utility.NowDateStr(),
		Year:           utility.NowYear(),
	}

	res, err := a.ar.PostActivity(attendance, tx)
	if err != nil {
		log.Printf("Failed to post end activity: %v", err)
		return nil, fmt.Errorf("failed to post end activity: %w", err)
	}

	return res, nil
}

// AddStartBreakTime 休憩の開始を登録
func (a ActivityServiceImpl) AddStartBreakTime(userID int, tx repository.Transaction) (*model.Attendance, error) {

	// 時間の登録
	attendance := &model.Attendance{
		UserID:         userID,
		AttendanceType: model.BreakStart,
		Time:           utility.NowTime(),
		Date:           utility.NowDateStr(),
		Year:           utility.NowYear(),
	}

	res, err := a.ar.PostActivity(attendance, tx)
	if err != nil {
		log.Printf("Failed to post end activity: %v", err)
		return nil, fmt.Errorf("failed to post end activity: %w", err)
	}

	return res, nil
}

// AddEndBreakTime 休憩の終了を登録
func (a ActivityServiceImpl) AddEndBreakTime(userID int, tx repository.Transaction) (*model.Attendance, error) {

	// 時間の登録
	attendance := &model.Attendance{
		UserID:         userID,
		AttendanceType: model.BreakEnd,
		Time:           utility.NowTime(),
		Date:           utility.NowDateStr(),
		Year:           utility.NowYear(),
	}

	res, err := a.ar.PostActivity(attendance, tx)
	if err != nil {
		log.Printf("Failed to post end activity: %v", err)
		return nil, fmt.Errorf("failed to post end activity: %w", err)
	}

	return res, nil
}

// EditTime 時間を修正
func (a ActivityServiceImpl) EditTime(activity *model.Attendance, newTime time.Time, tx repository.Transaction) (*model.Attendance, error) {

	// Dateからその日のユーザーの行動を全件取得
	dateStr := activity.Date
	userID := activity.UserID
	activityID := activity.ID
	historyByDate, err := a.hr.ReadHistoryByDate(userID, dateStr)
	if err != nil {
		return nil, err
	}

	// 編集時データのバリデーションチェック
	var beforeTime, afterTime time.Time

	// データが１つの
	for i, history := range historyByDate {
		if history.ID == activityID {

			// 最初の値かどうかを確認
			if i == 0 {
				if len(historyByDate) == 1 {
					beforeTime = utility.StartTime()
					afterTime = utility.NowTime()
					break
				}
				// 開始時刻を代入
				beforeTime = utility.StartTime()
				afterTime = historyByDate[i+1].Time
				break

				// 最後の値かどうか
			} else if i == len(historyByDate) {
				// 現在の日本時間を代入
				beforeTime = historyByDate[i-1].Time
				afterTime = utility.NowTime()
				break
			} else {
				beforeTime = historyByDate[i-1].Time
				afterTime = historyByDate[i+1].Time
				break
			}
		}
	}

	// 編集が前後の活動時間の間に含まれているかを確認
	if !utility.IsTimeInRange(beforeTime, newTime, afterTime) {
		return nil, utility.BadRequestError{}
	}

	// 編集するデータを詰め替える
	attendance := &model.Attendance{
		ID:   activityID,
		Time: newTime,
	}

	res, err := a.ar.PutActivity(attendance, tx)
	if err != nil {
		return nil, err
	}

	return res, nil
}

// Delete 活動を削除
func (a ActivityServiceImpl) Delete(attendance *model.Attendance, tx repository.Transaction) (*model.UserStatus, error) {
	dateStr := attendance.Date // 削除したい日の日付　"2023-12-25"
	userID := attendance.UserID
	activityID := attendance.ID
	nowDate := utility.NowDateStr()

	historyByDate, err := a.hr.ReadHistoryByDate(userID, dateStr)
	if err != nil {
		return nil, err
	}

	// 削除はその日の一番新しいものしかできないようにする(整合性を保つため)
	for i, history := range historyByDate {

		// 削除したいデータが日付ごとの最新のデータでない場合はエラーを返す
		if activityID == history.ID {
			if i != len(historyByDate)-1 {
				return nil, utility.BadRequestError{}
			}

			// 当日の削除の場合はユーザーの活動の状態も更新
			if dateStr == nowDate {

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
					return nil, err
				}
				// 変更後の状態を返す
				updateUserStatus := &model.UserStatus{
					UserID:   userID,
					StatusId: newAction,
				}
				// 削除処理を行う
				err = a.ar.DeleteActivity(activityID, tx)
				if err != nil {
					return nil, err
				}
				return updateUserStatus, nil
			} else {
				// 削除処理を行う
				err = a.ar.DeleteActivity(activityID, tx)
			}
		}
	}
	return nil, nil
}
