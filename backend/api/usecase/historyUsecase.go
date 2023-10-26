package usecase

import (
	"fmt"
	"github.com/yoshimi-I/AttendanceApp/domain/repository"
	"github.com/yoshimi-I/AttendanceApp/usecase/dto/response"
	"log"
)

// まずは扱う関数のinterfaceを実装
type HistoryUsecase interface {
	AllHistory(userId int, year int) ([]response.ActivityTimeResponseDTO, error)
	HistoryByDate(userId int, date string) (*response.HistoryByDateDto, error)
}

// 構造体を実装
type HistoryUsecaseImpl struct {
	hr repository.HistoryRepository
}

func (h HistoryUsecaseImpl) AllHistory(userId int, year int) ([]response.ActivityTimeResponseDTO, error) {
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
		date := v.Date
		attendanceType := v.AttendanceType
		duration := v.EndTime.Sub(v.StartTime)
		hours := int(duration.Hours()) //時間(h)切り捨て

		// 存在するかどうかのチェック
		if existDate, ok := durationMap[date]; ok {
			if attendanceType == 1 {
				durationMap[date] = existDate + hours
			} else {
				durationMap[date] = existDate - hours
			}
		} else {
			durationMap[date] = hours
		}
	}

	// DTOに変換
	for date, duration := range durationMap {
		responseData = append(responseData, response.ActivityTimeResponseDTO{
			Date:    date,
			SumTime: duration,
		})
	}
	return responseData, nil
}

func (h HistoryUsecaseImpl) HistoryByDate(userId int, date string) (*response.HistoryByDateDto, error) {
	historyByDate, err := h.hr.ReadHistoryByDate(userId, date)
	if err != nil {
		log.Println("failed to read historyByDate from HistoryRepository")
		return nil, err
	}

	var activities []response.ActivityDetail
	for _, v := range historyByDate {
		StartTime := v.StartTime
		EndTime := v.EndTime
		Type := response.ConvertActivityTime(v.AttendanceType)
		activity := response.ActivityDetail{
			Type:      Type,
			StartTime: StartTime,
			EndTime:   EndTime,
		}

		activities = append(activities, activity)
	}

	responseDto := &response.HistoryByDateDto{
		Date:       date,
		Activities: activities,
	}

	return responseDto, nil
}

// 関数を実装した構造体をnewする関数を実装,またこのとき返り値はinterface
func NewHistoryUsecase(hr repository.HistoryRepository) HistoryUsecase {
	return &HistoryUsecaseImpl{hr: hr}
}
