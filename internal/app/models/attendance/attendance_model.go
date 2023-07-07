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
