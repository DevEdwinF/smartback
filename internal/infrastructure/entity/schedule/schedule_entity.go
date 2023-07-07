package entity

import "time"

type ScheduleEntity struct {
	Id                      int    `json:"id" param:"id"`
	Day                     string `json:"day"`
	ArrivalTime             time.Time
	DepartureTime           time.Time
	FkCollaboratorsDocument int `json:"fk_collaborators_document"`
}
