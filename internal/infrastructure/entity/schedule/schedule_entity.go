package entity

type ScheduleEntity struct {
	Id        int    `json:"id" param:"id"`
	Arrival   string `json:"arrival"`
	Departure string `json:"departure"`
}
