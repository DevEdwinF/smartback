package models

type Schedules struct {
	Id            int `gorm:"primaryKey"`
	Day           string
	ArrivalTime   string
	DepartureTime string
	FkDocument    int
}
