package entity

// type AttendanceEntity struct {
// 	ID         int64  `json:"document"`
// 	Name       string `json:"name"`
// 	Position   string `json:"position"`
// 	ScheduleId int    `json:"schedule_id"`
// }

type AttendanceEntity struct {
	ID    int    `json:"document"`
	State string `json:"state"`
}

type ValidateSchedule struct {
	ID   string `json:"document"`
	Date string `json:"date"`
}
