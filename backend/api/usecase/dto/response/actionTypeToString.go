package response

import "log"

func ConvertActivityTime(actionType int) string {
	var Type string
	switch actionType {
	case 1:
		Type = "作業"
	case 2:
		Type = "休憩"
	case 3:
		Type = "その他"
	default:
		log.Printf("Invalid AttendanceType: %d", actionType)
	}
	return Type
}
