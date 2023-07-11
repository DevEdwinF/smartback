package entity

import "time"

<<<<<<< HEAD
type ScheduleEntity struct {
=======
type Schedule struct {
>>>>>>> a72ab65ff29686a4dc2d0ba391271818a54570b2
	Id                      int       `json:"id" param:"id"`
	Day                     string    `json:"day"`
	ArrivalTime             time.Time `json:"arrival_time"`
	DepartureTime           time.Time `json:"departure_time"`
	FkCollaboratorsDocument int       `json:"fk_collaborators_document"`
}
