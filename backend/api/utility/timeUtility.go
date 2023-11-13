package utility

import "time"

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

// IsTimeInRange　前後関係が正しいかどうかを確認する
func IsTimeInRange(startTime, targetTime, endTime time.Time) bool {
	return targetTime.After(startTime) && targetTime.Before(endTime)

}
