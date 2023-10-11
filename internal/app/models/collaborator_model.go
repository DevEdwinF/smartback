package models

import "time"

/* type Collaborators struct {
	Id             int
	Document       string
	FName          string
	LName          string
	Email          string
	Bmail          string
	State          string
	Leader         string
	LeaderDocument string
	Subprocess     string
	Headquarters   string
	Position       string
	CreatedAt      time.Time
} */

type Collaborators struct {
	Id             int    `gorm:"column:id" json:"id_collaborator"`
	Document       string `gorm:"column:document" json:"document"`
	FName          string `gorm:"column:f_name"`
	LName          string `gorm:"column:l_name"`
	Email          string `gorm:"column:email"`
	Bmail          string `gorm:"column:bmail"`
	State          string `gorm:"column:state"`
	Leader         string `gorm:"column:leader"`
	LeaderDocument string `gorm:"column:leader_document"`
	Subprocess     string `gorm:"column:subprocess"`
	Headquarters   string `gorm:"column:headquarters"`
	Position       string `gorm:"column:position"`
	CreatedAt      time.Time
}

// TableName specifies the table name for the mode
type CollaboratorsData struct {
	Collaborators
	Schedules
}

// type Employe struct {
// 	ID         int       `json:"id" gorm:"primary_key;auto_increment"`
// 	PinEmploye string    `json:"pinEmploye" gorm:"FOREIGNKEY:PinEmploye" `
// 	FirstName  string    `json:"first_name" `
// 	LastName   string    `json:"last_name"`
// 	Company    string    `json:"company"`
// 	Position   string    `json:"position"`
// 	ScheduleId int       `json:"schedule_id"`
// 	CreatedAt  time.Time `json:"fechacreacion"`
// }
