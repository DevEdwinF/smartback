package services

import (
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
)

func GetAllCollaborators() ([]entity.Collaborators, error) {
	collaboratorWithSchedule := []entity.Collaborators{}

	if err := config.DB.Table("collaborators").Select("*").Scan(&collaboratorWithSchedule).Error; err != nil {
		return nil, err
	}

	return collaboratorWithSchedule, nil
}
