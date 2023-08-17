package models

import (
	"database/sql"
	"time"
)

type Attendance struct {
	ID               int
	FkCollaboratorID int
	Location         string
	Arrival          sql.NullString
	Departure        sql.NullString
	Late             bool
	Photo            string
	CreatedAt        time.Time
}
type Translatedcollaborators struct {
	Id               int `json:"id"`
	FkCollaboratorId int
	FName            string `json:"f_name"`
	LName            string `json:"l_name"`
	// Document         string
	CreatedAt time.Time `json:"date"`
}

// func (Translated) TableName() string {
// 	return "translatedcollaborators"
// }
