package entity

import "time"

type ScheduleEntity struct {
	Id                      int       `json:"id" param:"id"`
	Day                     string    `json:"day"`
	ArrivalTime             time.Time `json:"arrival_time"`
	DepartureTime           time.Time `json:"departure_time"`
	FkCollaboratorsDocument int       `json:"fk_collaborators_document"`
}
