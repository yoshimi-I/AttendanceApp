package model

// ActionEnum ActionType を基に新しい型
type ActionEnum int

// 定数を定義し、iota を使用して連続した整数値を生成
const (
	WorkStart ActionEnum = iota + 1 // iotaは0から始まるので、+1して1から始める
	WorkEnd
	BreakStart
	BreakEnd
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
		return WorkStart
	case 2:
		return WorkEnd
	case 3:
		return BreakStart
	case 4:
		return BreakEnd
	case 5:
		return Pray
	default:
		return 0
	}
}

func (a ActionEnum) ToString() string {
	switch a {
	case WorkStart:
		return "WorkStart"
	case WorkEnd:
		return "WorkEnd"
	case BreakStart:
		return "BreakStart"
	case BreakEnd:
		return "BreakEnd"
	case Pray:
		return "Pray"
	default:
		return "Unknown"
	}
}
