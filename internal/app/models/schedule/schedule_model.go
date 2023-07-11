package models

import "time"

type ScheduleModel struct {
	Id            int
	Day           string
	ArrivalTime   time.Time
	DepartureTime time.Time
	FkDocument    int
}

func (Schedule) TableName() string {
	return "schedule"
}
