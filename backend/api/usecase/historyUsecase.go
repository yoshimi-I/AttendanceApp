package usecase

import (
	"fmt"
	"log"
	"time"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	"work-management-app/usecase/dto/response"
)

// まずは扱う関数のinterfaceを実装
type HistoryUsecase interface {
	AllHistory(userKey string, year int) ([]response.ActivityTimeResponseDTO, error)
	HistoryByDate(userKey string, date string) (*response.HistoryByDateDto, error)
}

// 構造体を実装
type HistoryUsecaseImpl struct {
	hr repository.HistoryRepository
	ur repository.UserRepository
}

// 関数を実装した構造体をnewする関数を実装,またこのとき返り値はinterface
func NewHistoryUsecase(hr repository.HistoryRepository, ur repository.UserRepository) HistoryUsecase {
	return &HistoryUsecaseImpl{
		hr: hr,
		ur: ur,
	}
}

func (h HistoryUsecaseImpl) AllHistory(userKey string, year int) ([]response.ActivityTimeResponseDTO, error) {
	// userKeyからuserIdを指定
	userId, err := h.ur.FindIDByUserKey(userKey)
	if err != nil {
		log.Println("usr_id not found")
		return nil, err
	}

	allHistoryList, err := h.hr.ReadAllHistory(userId, year)
	if err != nil {
		log.Println("failed to read all history from HistoryRepository")
		return nil, err
	}
	var responseData []response.ActivityTimeResponseDTO
	if err != nil {
		fmt.Println("no data")
	}

	// ここから合計時間のロジック
	var lastStartTime time.Time // 最後の作業開始時間
	var startDayTime time.Time  // 作業日の0:00:00

	durationMap := make(map[string]int)
	for _, v := range allHistoryList {

		attendanceType := v.AttendanceType

		activityTime := v.Time
		date := v.Date // 2023-11-12

		// まずはdateがkeyとして存在するかを確認
		if _, exist := durationMap[date]; !exist {
			durationMap[date] = 0
			// その日の0時0分0秒を取得
			startDayTime = time.Date(activityTime.Year(), activityTime.Month(), activityTime.Day(), 0, 0, 0, 0, activityTime.Location())

			// attendanceTypeが作業開始から始まってるかどうかを確認
			// 違っている場合は日付を跨いで何かをしている
			// 休憩開始,作業終了から記録が始まる場合は作業をしていたことになる
			// つまり0:00からの経過時間を足す必要がある
			if attendanceType == model.BreakStart || attendanceType == model.WorkEnd {
				duration := activityTime.Sub(startDayTime).Seconds() //経過時間を秒に変換
				durationMap[date] += int(duration)
			} else {
				// 作業開始から記録が始まっているので作業開始時間を保持
				lastStartTime = activityTime
			}

		} else {
			// dateの中で2回目以降のアクション
			// 休憩開始,作業終了をした時点で,作業時間をリセットして追加
			if attendanceType == model.BreakStart || attendanceType == model.WorkEnd {
				duration := activityTime.Sub(lastStartTime).Seconds()
				durationMap[date] += int(duration)
			} else {
				// 作業を開始した,または休憩を終了したので時間の計測を再開
				lastStartTime = activityTime
			}
		}
	}

	// DTOに変換
	for date, duration := range durationMap {
		durationHour := duration / 3600
		responseData = append(responseData, response.ActivityTimeResponseDTO{
			Year:         year,
			Date:         date,
			ActivityTime: durationHour,
		})
	}
	return responseData, nil
}

func (h HistoryUsecaseImpl) HistoryByDate(userKey string, date string) (*response.HistoryByDateDto, error) {
	// userKeyからuserIdを指定
	userId, err := h.ur.FindIDByUserKey(userKey)
	if err != nil {
		log.Println("usr_id not found")
		return nil, err
	}

	historyByDate, err := h.hr.ReadHistoryByDate(userId, date)
	if err != nil {
		log.Println("failed to read historyByDate from HistoryRepository")
		return nil, err
	}

	var activities []response.ActivityDetail
	for _, v := range historyByDate {
		Id := v.ID
		activityTime := v.Time
		activityType := response.ConvertActivityTime(v.AttendanceType)
		activity := response.ActivityDetail{
			Id:   Id,
			Type: activityType,
			Time: response.FormatChange(activityTime),
		}

		activities = append(activities, activity)
	}

	responseDto := &response.HistoryByDateDto{
		Date:       date,
		Activities: activities,
	}

	return responseDto, nil

}
