package model

// ActionEnum ActionType を基に新しい型
type ActionEnum int

// 定数を定義し、iota を使用して連続した整数値を生成
const (
	WorkStartEnd ActionEnum = iota + 1 // iotaは0から始まるので、+1して1から始める
	BreakStartEnd
	Pray
)

// ToInt ActionEnum 値を int 型に変換
func (a ActionEnum) ToInt() int {
	return int(a)
}

// IntToActionEnum は int 型を ActionEnum に変換
func IntToActionEnum(value int) ActionEnum {
	switch value {
	case 1:
		return WorkStartEnd
	case 2:
		return BreakStartEnd
	case 3:
		return Pray
	default:
		return 0
	}
}

func (a ActionEnum) ToString() string {
	switch a {
	case WorkStartEnd:
		return "WorkStartEnd"
	case BreakStartEnd:
		return "BreakStartEnd"
	case Pray:
		return "Pray"
	default:
		return "Unknown"
	}
}
