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
	Id               int
	FkCollaboratorId int
	// Document         string
	CreatedAt time.Time
}

// func (Translated) TableName() string {
// 	return "translatedcollaborators"
// }
