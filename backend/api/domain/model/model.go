package model

import "time"

type Users struct {
	UserId int
	Name   string
	Email  string
}

type StudyHistory struct {
	StudyTime time.Time
	StudyDay  time.Time
}

type Play struct {
	ID        int
	StartTime time.Time
	EndTime   time.Time
}

type StudyStartEnd struct {
	ID        int
	StartTime time.Time
	EndTime   time.Time
}

type BreakStartEnd struct {
	ID        int
	StartTime time.Time
	EndTime   time.Time
}
