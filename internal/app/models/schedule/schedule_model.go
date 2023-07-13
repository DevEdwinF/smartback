package models

type Schedule struct {
	Id            int `gorm:"primaryKey"`
	Day           string
	ArrivalTime   string
	DepartureTime string
	FkDocument    int
}

func (Schedule) TableName() string {
	return "schedule"
}
