package models

type Schedule struct {
	Id        int    `json:"id_sch" gorm:"primary_key;auto_increment"`
	Arrival   string `json:"arrival"`
	Departure string `json:"departure"`
}
