package orm_model

import (
	"time"
)

type User struct {
	Id        int       `gorm:"primaryKey;column:id"`
	Name      string    `gorm:"column:name"`
	Email     string    `gorm:"column:email"`
	UserKey   string    `gorm:"column:user_key"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type AttendanceType struct {
	AttendanceTypeId int       `gorm:"primaryKey;column:attendance_type_id"`
	ActionType       string    `gorm:"column:action_type"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}

type Attendance struct {
	Id             int       `gorm:"primaryKey;column:id"`
	UserId         int       `gorm:"column:user_id"`
	AttendanceType int       `gorm:"column:attendance_type"`
	Time           time.Time `gorm:"column:time"`
	Date           string    `gorm:"column:date"`
	Year           int       `gorm:"column:year"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}

type UserStatus struct {
	UserId    int       `gorm:"primaryKey;column:user_id"`
	StatusId  int       `gorm:"column:status_id"`
	CreatedAt time.Time `gorm:"column:created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at"`
}

type UserStatusType struct {
	UserStatusTypeId int       `gorm:"primaryKey;column:user_status_type_id"`
	UserStatusType   int       `gorm:"column::user_status_id"`
	CreatedAt        time.Time `gorm:"column:created_at"`
	UpdatedAt        time.Time `gorm:"column:updated_at"`
}
