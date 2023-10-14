package orm

import (
	"time"
)

// Users corresponds to the Users table in the database.
type Users struct {
	UserID     uint         `gorm:"primaryKey;autoIncrement"`
	Name       string       `gorm:"type:varchar(255)"`
	Email      string       `gorm:"type:varchar(255);unique"`
	CreatedAt  time.Time    `gorm:"autoCreateTime"`
	UpdatedAt  time.Time    `gorm:"autoUpdateTime"`
	Activities []Activities `gorm:"foreignKey:UserID"`
}

// Activities corresponds to the Activities table in the database.
type Activities struct {
	ActivityID   uint `gorm:"primaryKey;autoIncrement"`
	UserID       uint
	ActivityDate time.Time
	Notes        string          `gorm:"type:text"`
	CreatedAt    time.Time       `gorm:"autoCreateTime"`
	UpdatedAt    time.Time       `gorm:"autoUpdateTime"`
	Plays        []Play          `gorm:"foreignKey:ActivityID"`
	Studies      []StudyStartEnd `gorm:"foreignKey:ActivityID"`
	Breaks       []BreakStartEnd `gorm:"foreignKey:ActivityID"`
}

// Play corresponds to the Play table in the database.
type Play struct {
	PlayID     uint `gorm:"primaryKey;autoIncrement"`
	ActivityID uint
	StartTime  time.Time
	EndTime    time.Time
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

// StudyStartEnd corresponds to the StudyStartEnd table in the database.
type StudyStartEnd struct {
	StudyID    uint `gorm:"primaryKey;autoIncrement"`
	ActivityID uint
	StartTime  time.Time
	EndTime    time.Time
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}

// BreakStartEnd corresponds to the BreakStartEnd table in the database.
type BreakStartEnd struct {
	BreakID    uint `gorm:"primaryKey;autoIncrement"`
	ActivityID uint
	StartTime  time.Time
	EndTime    time.Time
	CreatedAt  time.Time `gorm:"autoCreateTime"`
	UpdatedAt  time.Time `gorm:"autoUpdateTime"`
}
