package entity

type AttendanceEntity struct {
	ID         int64  `json:"document"`
	Name       string `json:"name"`
	Position   string `json:"position"`
	ScheduleId int    `json:"schedule_id"`
}
