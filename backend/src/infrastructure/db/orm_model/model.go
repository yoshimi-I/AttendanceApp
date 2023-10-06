package orm_model

import (
	"time"
)

type Employee struct {
	EmployeeID uint `gorm:"primaryKey"`
	Name       string
	Attendances []Attendance `gorm:"foreignKey:EmployeeID"`
}

type AttendanceStatus struct {
	StatusID   uint   `gorm:"primaryKey"`
	StatusName string
	Attendances []Attendance `gorm:"foreignKey:StatusID"`
}

type Attendance struct {
	AttendanceID   uint      `gorm:"primaryKey"`
	EmployeeID     uint
	AttendanceDate time.Time
	ClockTime      time.Time
	StatusID       uint
	Notes          string
}
