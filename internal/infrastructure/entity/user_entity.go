package entity

import "time"

type User struct {
	Document string `json:"document"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CollaboratorsEntity struct {
	Document int       `json:"document"`
	FName    string    `json:"f_name"`
	LName    string    `json:"l_name"`
	Email    string    `json:"email"`
	Position string    `json:"position"`
	Leader   string    `json:"leader"`
	CreateAt time.Time `json:"date"`
}

type CollaboratorsDataEntity struct {
	CollaboratorsEntity
	Schedule
}
