package model

import "time"

type Users struct {
	UserId int
	Name   string
	Email  string
}

type Play struct {
	PlayID     int
	ActivityID int
	StartTime  time.Time
	EndTime    time.Time
}

type StudyStartEnd struct {
	StudyID    int
	ActivityID int
	StartTime  time.Time
	EndTime    time.Time
}

type BreakStartEnd struct {
	BreakID    int
	ActivityID int
	StartTime  time.Time
	EndTime    time.Time
}

type Activities struct {
	ActivityDate time.Time
	Notes        string
	Plays        []Play
	Studies      []StudyStartEnd
	Breaks       []BreakStartEnd
	SumTime      int
}
