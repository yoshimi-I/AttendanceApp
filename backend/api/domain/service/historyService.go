package service

import (
	"time"
	"work-management-app/domain/model"
	"work-management-app/utility"
)

type HistoryDomainService interface {
	TotalWorkByYear(allHistoryList []model.Attendance) map[string]int
}

type HistoryServiceImpl struct{}

func NewHistoryDomainService() HistoryDomainService {
	return &HistoryServiceImpl{}
}

// TotalWorkByYear 1日の合計活動時間をカウント
func (h HistoryServiceImpl) TotalWorkByYear(allHistoryList []model.Attendance) map[string]int {

	var lastStartTime time.Time // 最後の作業開始時間
	var startDayTime time.Time  // 作業日の0:00:00

	durationMap := make(map[string]int)
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
	for date, duration := range durationMap {
		durationHour := duration / 3600

		// 1h切ってても作業があれば1にする
		if durationHour == 0 {
			durationHour += 1
		}
		durationMap[date] = durationHour
	}

	return durationMap
}
