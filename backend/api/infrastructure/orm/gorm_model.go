package orm_model

import (
	"time"
)

type User struct {
	ID        int       `gorm:"primaryKey;column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	UserKey   string    `gorm:"column:user_key"`
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
	Time           time.Time `gorm:"column:time"`
	Date           string    `gorm:"column:date"`
	Year           int       `gorm:"column:year"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

type UserStatus struct {
	UserID    int       `gorm:"primaryKey;column:user_id"`
	StatusID  int       `gorm:"column:status_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type UserStatusType struct {
	UserStatusTypeID int       `gorm:"primaryKey;column:user_status_type_id"`
	UserStatusType   int       `gorm:"column::user_status_id"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}
