package entity

type Schedules struct {
	Id               int    `json:"id" param:"id"`
	Day              string `json:"day"`
	ArrivalTime      string `json:"arrival_time"`
	DepartureTime    string `json:"departure_time"`
	FkCollaboratorId int    `json:"fk_collaborator_id"`
	// Document         *string `json:"document,omitempty"`
}
