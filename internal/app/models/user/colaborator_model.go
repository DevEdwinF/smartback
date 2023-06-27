package models

import (
	"time"

	models "github.com/DevEdwinF/smartback.git/internal/app/models/attendance"
)

type Colaborators struct {
	ID       int
	document string
	Name     string
	email    string
	CreateAt time.Time
}

type EmployeeWithSchedule struct {
	Colaborators
	models.Schedule
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
