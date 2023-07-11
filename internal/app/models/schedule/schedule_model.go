package models

import "time"

type Schedule struct {
	Id            int `gorm:"primaryKey"`
	Day           string
	ArrivalTime   time.Time
	DepartureTime time.Time
	FkDocument    int
}

func (Schedule) TableName() string {
	return "schedule"
}
