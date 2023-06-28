package models

import (
	"time"
)

type Attendance struct {
	ID           int
	FkDocumentId int
	Arrival      *time.Time
	Departure    *time.Time
	CreatedAt    time.Time
}
