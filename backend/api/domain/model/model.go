package model

import "time"

type User struct {
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

type Activity struct {
	Date          time.Time
	Plays         []Play
	StudySessions []StudyStartEnd
	Breaks        []BreakStartEnd
}
