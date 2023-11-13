package utility

import (
	"fmt"
	"time"
)

// StartTime 始まり（0:00）の日本時間を返す
func StartTime() time.Time {
	now := time.Now()
	jst, _ := time.LoadLocation("Asia/Tokyo")
	start := time.Date(now.Year(), now.Month(), now.Day(), 0, 0, 0, 0, jst)
	return start
}

// EndTime 終わり（23:59:59）の日本時間を返す
func EndTime() time.Time {
	now := time.Now()
	jst, _ := time.LoadLocation("Asia/Tokyo")
	end := time.Date(now.Year(), now.Month(), now.Day(), 23, 59, 59, 0, jst)
	return end
}

// NowTime 現在の日本時間を返す
func NowTime() time.Time {
	now := time.Now()
	jst, _ := time.LoadLocation("Asia/Tokyo")
	return now.In(jst)
}

// NowDateStr 現在の日付を文字列型で返す ("2023-12-25")
func NowDateStr() string {
	now := time.Now()
	jst, _ := time.LoadLocation("Asia/Tokyo")
	nowStr := now.In(jst)
	return fmt.Sprintf("%d-%02d-%02d", nowStr.Year(), nowStr.Month(), nowStr.Day())
}

// PreDateStr 前日の日付を文字列で返す
func PreDateStr(nowDate string) (string, error) {
	const layout = "2006-01-02"
	date, err := time.Parse(layout, nowDate)

	if err != nil {
		return "", err
	}
	
	previousDay := date.AddDate(0, 0, -1)
	return previousDay.Format(layout), err
}

// IsTimeInRange　前後関係が正しいかどうかを確認する
func IsTimeInRange(startTime, targetTime, endTime time.Time) bool {
	return targetTime.After(startTime) && targetTime.Before(endTime)

}
