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

/*
	 func (c *CollaboratorFilter) BuildQuery() string {
		query := ""
		if c.Document != "" {
			query += " OR document = " + c.Document
		}
		if c.FName != "" {
			query += " OR f_name = " + c.FName
		}
		if c.LName != "" {
			query += " OR l_name = " + c.LName
		}
		if c.Bmail != "" {
			query += " OR bmail = " + c.Bmail
		}
		if c.Email != "" {
			query += " OR email = " + c.Email
		}
		if c.Position != "" {
			query += " OR position = " + c.Position
		}
		if c.State != "" {
			query += " OR state = " + c.State
		}
		if c.Leader != "" {
			query += " OR leader = " + c.Leader
		}
		if c.LeaderDocument != "" {
			query += " OR leader_document = " + c.LeaderDocument
		}
		if c.Subprocess != "" {
			query += " OR subprocess = " + c.Subprocess
		}
		if c.Headquarters != "" {
			query += " OR headquarters = " + c.Headquarters
		}
		return query
	}
*/
type CollaboratorsDataEntity struct {
	Collaborators
	Schedules
}
