package models

import "time"

type Schedule struct {
	Id            int
	Day           string
	ArrivalTime   time.Time
	DepartureTime time.Time
	FkDocument    int
}
