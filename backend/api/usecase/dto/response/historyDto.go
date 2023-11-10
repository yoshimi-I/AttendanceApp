package response

type AllHistoryDto struct {
	Day      string `json:"day"`
	Activity int    `json:"activity"`
}

type ActivityDetail struct {
	Id   int    `json:"id"`
	Type string `json:"type"`
	Time string `json:"time"`
}

type HistoryByDateDto struct {
	Date       string           `json:"date"`
	Activities []ActivityDetail `json:"activities"`
}

func ConvertAllHistory(day string, activity int) *AllHistoryDto {
	return &AllHistoryDto{
		Day:      day,
		Activity: activity,
	}
}
