package entity

type Schedule struct {
	Id                      int    `json:"id" param:"id"`
	Day                     string `json:"day"`
	ArrivalTime             string `json:"arrival_time"`
	DepartureTime           string `json:"departure_time"`
	FkCollaboratorsDocument int    `json:"fk_collaborators_document"`
}
