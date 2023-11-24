package usecase

import (
	"fmt"
	"log"
	"work-management-app/application/dto/request"
	"work-management-app/application/dto/response"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	"work-management-app/domain/service"
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
	db repository.Transaction
	ar repository.ActivityRepository
	ur repository.UserRepository
	as service.ActivityDomainService
}

func NewActivityUsecase(ar repository.ActivityRepository, ur repository.UserRepository, as service.ActivityDomainService, db repository.Transaction) ActivityUsecase {
	return &ActivityUsecaseImpl{
		ar: ar,
		ur: ur,
		as: as,
		db: db,
	}
}

// AddStarWork　作業の開始を登録
func (a ActivityUsecaseImpl) AddStarWork(work *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error) {

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
		return nil, fmt.Errorf("you can't start work")
	}

	// トランザクション開始
	tx, err := a.db.TxBegin()
	if err != nil {
		return nil, err
	}

	// エラーハンドリング用のdefer
	defer func() {
		if err != nil {
			err := a.db.TxRollback()
			if err != nil {
				log.Panic(err)
			}
		} else {
			err := a.db.TxCommit()
			if err != nil {
				log.Panic(err)
			}
		}
	}()

	// 作業開始を登録
	res, err := a.as.AddStarWorkTime(userID, tx)
	if err != nil {
		return nil, err
	}

	// ユーザーの状態の更新
	newUserStatus := &model.UserStatus{
		UserID:   userID,
		StatusId: model.Work,
	}

	_, err = a.ur.PutUserStatus(newUserStatus, tx)
	if err != nil {
		return nil, err
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		Id:             res.ID,
		AttendanceType: model.WorkStart.ToString(),
		Time:           res.Time,
		Year:           res.Year,
		Date:           res.Date,
		Status:         model.Work.ToString(),
	}
	return responseDTO, nil
}

// AddEndWork 作業の終了を登録
func (a ActivityUsecaseImpl) AddEndWork(work *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error) {
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
		return nil, fmt.Errorf("you can't end work")
	}

	// トランザクション開始
	tx, err := a.db.TxBegin()
	if err != nil {
		return nil, err
	}

	// エラーハンドリング用のdefer
	defer func() {
		if err != nil {
			err := a.db.TxRollback()
			if err != nil {
				log.Panic(err)
			}
		} else {
			err := a.db.TxCommit()
			if err != nil {
				log.Panic(err)
			}
		}
	}()

	// ユーザーの作業終了を登録
	res, err := a.as.AddEndWorkTime(userID, tx)
	if err != nil {
		return nil, err
	}

	// ユーザーの状態の更新
	newUserStatus := &model.UserStatus{
		UserID:   userID,
		StatusId: model.Finish,
	}

	_, err = a.ur.PutUserStatus(newUserStatus, tx)
	if err != nil {
		return nil, err
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		Id:             res.ID,
		AttendanceType: model.WorkEnd.ToString(),
		Time:           res.Time,
		Year:           res.Year,
		Date:           res.Date,
		Status:         model.Finish.ToString(),
	}
	return responseDTO, nil
}

// AddStartBreak　休憩の開始を登録
func (a ActivityUsecaseImpl) AddStartBreak(breakInfo *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error) {

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
		return nil, fmt.Errorf("you can't start break")
	}

	if err != nil {
		return nil, err
	}

	// トランザクション開始
	tx, err := a.db.TxBegin()
	if err != nil {
		return nil, err
	}

	// エラーハンドリング用のdefer
	defer func() {
		if err != nil {
			err := a.db.TxRollback()
			if err != nil {
				log.Panic(err)
			}
		} else {
			err := a.db.TxCommit()
			if err != nil {
				log.Panic(err)
			}
		}
	}()
	// 休憩開始を登録
	res, err := a.as.AddStartBreakTime(userID, tx)
	if err != nil {
		log.Printf("you can't start break")
		return nil, fmt.Errorf("failed to post start break: %w", err)
	}

	// ユーザーの状態の更新
	newUserStatus := &model.UserStatus{
		UserID:   userID,
		StatusId: model.Break,
	}

	_, err = a.ur.PutUserStatus(newUserStatus, tx)
	if err != nil {
		return nil, err
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		Id:             res.ID,
		AttendanceType: model.BreakStart.ToString(),
		Time:           res.Time,
		Year:           res.Year,
		Date:           res.Date,
		Status:         model.Break.ToString(),
	}
	return responseDTO, nil
}

// AddEndBreak 休憩の終了を登録
func (a ActivityUsecaseImpl) AddEndBreak(breakInfo *request.ActivityRequestDTO) (*response.ActivityResponseDTO, error) {

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
		return nil, fmt.Errorf("you can't end break")
	}

	// トランザクション開始
	tx, err := a.db.TxBegin()
	if err != nil {
		return nil, err
	}

	// エラーハンドリング用のdefer
	defer func() {
		if err != nil {
			err := a.db.TxRollback()
			if err != nil {
				log.Panic(err)
			}
		} else {
			err := a.db.TxCommit()
			if err != nil {
				log.Panic(err)
			}
		}
	}()

	// ユーザーの作業終了を登録
	res, err := a.as.AddEndBreakTime(userID, tx)
	if err != nil {
		log.Printf("Failed to post end break: %v", err)
		return nil, fmt.Errorf("failed to post end break: %w", err)
	}

	// ユーザーの状態の更新
	newUserStatus := &model.UserStatus{
		UserID:   userID,
		StatusId: model.Work,
	}

	_, err = a.ur.PutUserStatus(newUserStatus, tx)
	if err != nil {
		return nil, err
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		Id:             res.ID,
		AttendanceType: model.BreakEnd.ToString(),
		Time:           res.Time,
		Year:           res.Year,
		Date:           res.Date,
		Status:         model.Work.ToString(),
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
		return nil, utility.NotFoundError{}
	}

	// 編集処理をする人が本人かどうかを確認
	if userID != record.UserID {
		return nil, utility.AuthenticationError{}
	}

	// トランザクション開始
	tx, err := a.db.TxBegin()
	if err != nil {
		return nil, err
	}

	// エラーハンドリング用のdefer
	defer func() {
		if err != nil {
			err := a.db.TxRollback()
			if err != nil {
				log.Panic(err)
			}
		} else {
			err := a.db.TxCommit()
			if err != nil {
				log.Panic(err)
			}
		}
	}()

	// データの編集を行う
	res, err := a.as.EditTime(record, newTime, tx)
	if err != nil {
		return nil, err
	}

	// DTOに詰め替え作業
	responseDTO := &response.ActivityResponseDTO{
		Id:             res.ID,
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
		return utility.NotFoundError{}
	}

	// 編集処理をする人の権限確認
	if userID != record.UserID {
		return utility.AuthenticationError{}
	}

	// トランザクション開始
	tx, err := a.db.TxBegin()
	if err != nil {
		log.Panic(err)
	}

	// エラーハンドリング用のdefer
	defer func() {
		if err != nil {
			err := a.db.TxRollback()
			if err != nil {
				log.Panic(err)
			}
		} else {
			err := a.db.TxCommit()
			if err != nil {
				log.Panic(err)
			}
		}
	}()

	// 削除
	newUserStatus, err := a.as.Delete(record, tx)
	if err != nil {
		return err
	}

	// ユーザーの状態の更新
	if newUserStatus != nil {
		_, err = a.ur.PutUserStatus(newUserStatus, tx)
		if err != nil {
			return err
		}
	}

	return nil

}
