package usecase

import (
	"fmt"
	"log"
	"time"
	"work-management-app/domain/model"
	"work-management-app/domain/repository"
	"work-management-app/usecase/dto/response"
	"work-management-app/utility"
)

// HistoryUsecase まずは扱う関数のinterfaceを実装
type HistoryUsecase interface {
	AllHistory(userKey string, year int) ([]response.ActivityTimeResponseDTO, error)
	HistoryByDate(userKey string, date string) (*response.HistoryByDateDto, error)
}

// HistoryUsecaseImpl 構造体を実装
type HistoryUsecaseImpl struct {
	hr repository.HistoryRepository
	ur repository.UserRepository
}

// NewHistoryUsecase 関数を実装した構造体をnewする関数を実装,またこのとき返り値はinterface
func NewHistoryUsecase(hr repository.HistoryRepository, ur repository.UserRepository) HistoryUsecase {
	return &HistoryUsecaseImpl{
		hr: hr,
		ur: ur,
	}
}

// AllHistory 年を指定してその活動を取得
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

	// ここから合計時間を計算する処理
	var lastStartTime time.Time // 最後の作業開始時間
	var startDayTime time.Time  // 作業日の0:00:00

	durationMap := make(map[string]int)
	fmt.Println(allHistoryList)
	for _, v := range allHistoryList {

		attendanceType := v.AttendanceType

		activityTime := v.Time
		date := v.Date // 2023-11-12

		// まずはdateが配列の中に存在するかを確認
		if _, exist := durationMap[date]; !exist {
			durationMap[date] = 0
			// その日の0時0分0秒を取得
			startDayTime = utility.StartTime()

			// attendanceTypeが作業開始から始まってるかどうかを確認
			// 休憩開始,作業終了から記録が始まる場合は日付を跨いで作業している
			// つまり0:00からの経過時間を足す必要がある
			if attendanceType == model.BreakStart || attendanceType == model.WorkEnd {
				duration := activityTime.Sub(startDayTime).Seconds() //経過時間を秒に変換
				durationMap[date] += int(duration)

				// さらに前の日の作業にも0:00までの時間を追加する必要がある
				// 前日の最終時間を取得
				endDayTime := startDayTime.Add(-time.Second)
				duration = endDayTime.Sub(lastStartTime).Seconds() //経過時間を秒に変換

				// ここから24hを引いていき,durationが無くなるまで前の日に入れていく
				daySumSecond := float64(24 * 60 * 60) // 1日の合計(秒)
				for duration >= 0 {
					println(duration)
					// 一つ前の日付を取得
					preDate, _ := utility.PreDateStr(date)
					if duration >= daySumSecond {
						durationMap[preDate] = int(daySumSecond)
					} else {
						durationMap[preDate] = int(duration)
					}
					duration -= daySumSecond
					date = preDate
				}

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

		// 1h切ってても作業があれば1にする
		if durationHour == 0 {
			durationHour += 1
		}
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
		ID := v.ID
		activityTime := v.Time
		activityType := response.ConvertActivityTime(v.AttendanceType)
		activity := response.ActivityDetail{
			Id:   ID,
			Type: activityType,
			Time: activityTime,
		}

		activities = append(activities, activity)
	}

	responseDto := &response.HistoryByDateDto{
		Date:       date,
		Activities: activities,
	}

	return responseDto, nil

}
