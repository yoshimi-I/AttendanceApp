package usecase

import (
	"log"
	"work-management-app/application/dto/response"
	"work-management-app/domain/repository"
	"work-management-app/domain/service"
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
	hs service.HistoryDomainService
}

// NewHistoryUsecase 関数を実装した構造体をnewする関数を実装,またこのとき返り値はinterface
func NewHistoryUsecase(
	hr repository.HistoryRepository,
	ur repository.UserRepository,
	hs service.HistoryDomainService,
) HistoryUsecase {
	return &HistoryUsecaseImpl{
		hr: hr,
		ur: ur,
		hs: hs,
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

	// 年を指定してその年のすべての活動を取得
	allHistoryList, err := h.hr.ReadAllHistory(userId, year)
	if err != nil {
		log.Println("failed to read all history from HistoryRepository")
		return nil, err
	}

	// 合計時間を計算する処理
	durationMap := h.hs.TotalWorkByYear(allHistoryList)

	// DTOに変換
	var responseData []response.ActivityTimeResponseDTO
	for date, duration := range durationMap {
		responseData = append(responseData, response.ActivityTimeResponseDTO{
			Year:         year,
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

	// DTOに詰め替える
	responseDto := &response.HistoryByDateDto{
		Date:       date,
		Activities: activities,
	}

	return responseDto, nil

}
