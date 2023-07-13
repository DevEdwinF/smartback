package models

type Schedule struct {
	Id            int    `json:"id_sch" gorm:"primary_key;auto_increment"`
	ArrivalTime   string `json:"arrival"`
	DepartureTime string `json:"departure"`
}

func (s *Schedule) TableName() string {
	return "schedule"
}
