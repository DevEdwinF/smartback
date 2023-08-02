package services

import (
	"errors"
	"fmt"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"gorm.io/gorm"
)

func GetAllCollaborators() ([]entity.Collaborators, error) {
	collaboratorWithSchedule := []entity.Collaborators{}

	if err := config.DB.Table("collaborators").Select("*").Scan(&collaboratorWithSchedule).Error; err != nil {
		return nil, err
	}

	return collaboratorWithSchedule, nil
}

func ValidateCollaboratorService(document string) (*models.Collaborators, error) {
	var collaborator models.Collaborators
	err := config.DB.Model(&collaborator).Where("document = ?", document).First(&collaborator).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, fmt.Errorf("Colaborador no encontrado")
		}
		return nil, err
	}
	return &collaborator, nil
}
