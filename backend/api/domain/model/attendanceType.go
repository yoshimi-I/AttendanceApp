package model

// ActionType を基に新しい型 ActionEnum を定義
type ActionEnum int

// 定数を定義し、iota を使用して連続した整数値を生成
const (
	WorkStartEnd ActionEnum = iota
	Break
	Pray
)

// ActionType の文字列表現を返す関数
func (a ActionEnum) String() string {
	return [...]string{"WorkStartEnd", "Break", "Pray"}[a]
}

type AttendanceType struct {
	AttendanceTypeID int
	ActionType       ActionEnum
}
