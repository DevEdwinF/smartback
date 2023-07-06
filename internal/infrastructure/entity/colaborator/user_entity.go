package entity

import (
	"time"

	entity "github.com/DevEdwinF/smartback.git/internal/infrastructure/entity/schedule"
)

type User struct {
	Document string `json:"document"`
	Email    string `json:"email"`
	Pass     string `json:"pass"`
}

type CollaboratorsEntity struct {
	Document   int       `json:"document"`
	Name       string    `json:"name"`
	Email      string    `json:"email"`
	Position   string    `json:"position"`
	ScheduleId int       `json:"fk_schedule_id"`
	CreateAt   time.Time `json:"date"`
}

type CollaboratorsDataEntity struct {
	CollaboratorsEntity
	entity.ScheduleEntity
}
