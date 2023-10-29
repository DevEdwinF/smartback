package entity

import "time"

type User struct {
	Document string `json:"document"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserData struct {
	User
	FName     string    `json:"f_name"`
	LName     string    `json:"l_name"`
	CreatedAt time.Time `json:"created_at"`
	FkRoleId  int       `json:"rol"`
}

type UserFilter struct {
	UserData
	Paginate
	RoleName string
}

type Users []UserData

type Collaborators struct {
	ID             int       `json:"id_collaborator" query:"id_collaborator"`
	Document       string    `json:"document" query:"document"`
	FName          string    `json:"f_name" query:"f_name"`
	LName          string    `json:"l_name" query:"l_name"`
	Bmail          string    `json:"bmail" query:"bmail"`
	Email          string    `json:"email" query:"email"`
	Position       string    `json:"position" query:"position"`
	State          string    `json:"state" query:"state"`
	Leader         string    `json:"leader" query:"leader"`
	LeaderDocument string    `json:"leader_document" query:"leader_document"`
	Subprocess     string    `json:"subprocess" query:"subprocess"`
	Headquarters   string    `json:"headquarters" query:"headquarters"`
	CreatedAt      time.Time `json:"date" query:"date"`
}

type CollaboratorFilter struct {
	Collaborators
	Paginate
	LeaderDocument string
}

type CollaboratorsDataEntity struct {
	Collaborators
	Schedules
}
