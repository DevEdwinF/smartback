package models

import (
	"time"

	models "github.com/DevEdwinF/smartback.git/internal/app/models/schedule"
)

type Collaborators struct {
	// ID       int
	Document int
	Name     string
	Email    string
	Position string
	CreateAt time.Time
}

type CollaboratorsData struct {
	Collaborators
	models.ScheduleModel
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
