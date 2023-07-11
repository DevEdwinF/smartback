package models

import (
	"time"
)

type Attendance struct {
	ID           int
	FkDocumentId int
	Location     string
	Arrival      *time.Time
	Departure    *time.Time
	Photo        string
	CreatedAt    time.Time
}

type translatedcollaborators struct {
	ID           int
	FkDocumentId int
	CreatedAt    time.Time
}

// func (Translated) TableName() string {
// 	return "translatedcollaborators"
// }
