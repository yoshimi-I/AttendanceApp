package model

// StatusEnum を基に新しい型を定義
type StatusEnum int

// 定数を定義し、iota を使用して連続した整数値を生成
const (
	Work StatusEnum = iota + 1 // iotaは0から始まるので、+1して1から始める
	Break
	Finish
)

func (e StatusEnum) ToInt() int {
	return int(e)
}

func IntToStatusEnum(value int) StatusEnum {
	switch value {
	case 1:
		return Work
	case 2:
		return Break
	case 3:
		return Finish
	default:
		return 0
	}
}

func (e StatusEnum) ToString() string {
	switch e {
	case Work:
		return "Work"
	case Break:
		return "Break"
	case Finish:
		return "Finish"
	default:
		return "Unknown"
	}
}
