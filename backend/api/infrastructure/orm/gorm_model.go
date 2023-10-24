package orm_model

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type AttendanceType struct {
	AttendanceTypeID int       `gorm:"primaryKey;column:attendance_type_id"`
	ActionType       string    `gorm:"column:action_type"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

type Attendance struct {
	ID             int       `gorm:"primaryKey;column:id"`
	UserID         int       `gorm:"column:user_id"`
	AttendanceType int       `gorm:"column:attendance_type"`
	StartTime      time.Time `gorm:"column:start_time"`
	EndTime        time.Time `gorm:"column:end_time"`
	Date           string    `gorm:"column:date"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}
