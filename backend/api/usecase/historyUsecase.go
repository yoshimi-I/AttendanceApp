package usecase

import (
	"fmt"
	"log"
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

	durationMap := make(map[string]int)

	for _, v := range allHistoryList {
		date := v.Date                     // yyyy-mm-dd
		attendanceType := v.AttendanceType // 作業か,休憩かを判断
		duration := v.EndTime.Sub(v.StartTime)

		hours := int(duration.Minutes()) / 60 // 分(h)切り捨て.その後60で割る
		// 存在するかどうかのチェック
		if _, ok := durationMap[date]; ok {
			if attendanceType == 1 {
				durationMap[date] += hours
			} else {
				durationMap[date] -= hours
			}
		} else {
			durationMap[date] = hours
		}
	}

	// DTOに変換
	for date, duration := range durationMap {
		responseData = append(responseData, response.ActivityTimeResponseDTO{
			Date:         date,
			ActivityTime: duration,
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
		StartTime := v.StartTime
		EndTime := v.EndTime
		Type := response.ConvertActivityTime(v.AttendanceType)
		activity := response.ActivityDetail{
			Id:        Id,
			Type:      Type,
			StartTime: response.FormatChange(StartTime),
			EndTime:   response.FormatChange(EndTime),
		}

		activities = append(activities, activity)
	}

	responseDto := &response.HistoryByDateDto{
		Date:       date,
		Activities: activities,
	}

	return responseDto, nil

}
