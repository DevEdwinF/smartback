package entity

import "time"

// type AttendanceEntity struct {
// 	ID         int64  `json:"document"`
// 	Name       string `json:"name"`
// 	Position   string `json:"position"`
// 	ScheduleId int    `json:"schedule_id"`
// }

type AttendanceEntity struct {
	FkDocumentId string `json:"document"`
	// Document  string    `json:"document"`
	State     string    `json:"state"`
	CreatedAt time.Time `json:"date"`
}

type ValidateSchedule struct {
	Id   string `json:"document"`
	Date string `json:"date"`
}
