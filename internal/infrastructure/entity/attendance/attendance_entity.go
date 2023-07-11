package entity

import "time"

// type AttendanceEntity struct {
// 	ID         int64  `json:"document"`
// 	Name       string `json:"name"`
// 	Position   string `json:"position"`
// 	ScheduleId int    `json:"schedule_id"`
// }

type AttendanceEntity struct {
	FkDocumentId int       `json:"document"`
	State        string    `json:"state"`
	Location     string    `josn:"location"`
	Photo        string    `json:"photo"`
	CreatedAt    time.Time `json:"date"`
}

type UserAttendanceData struct {
	FkDocumentId int        `json:"document"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Location     string     `josn:"location"`
	Arrival      *time.Time `json:"arrival"`
	Departure    *time.Time `json:"departure"`
	Photo        string     `json:"photo"`
	CreatedAt    time.Time  `json:"date"`
}

type ValidateSchedule struct {
	Id   string `json:"document"`
	Date string `json:"date"`
}

type Translatedcollaborators struct {
	FkDocumentId int       `json:"document"`
	CreatedAt    time.Time `json:"date"`
}
