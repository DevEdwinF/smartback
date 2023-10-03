package services

import (
	"errors"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"gorm.io/gorm"
)

type CollaboratorService struct {
	db *gorm.DB
}

func NewCollaboratorService(db *gorm.DB) *CollaboratorService {
	return &CollaboratorService{db: db}
}

func GetAllCollaborators() ([]entity.Collaborators, error) {
	collaboratorWithSchedule := []entity.Collaborators{}

	if err := config.DB.Table("collaborators").
		Select("*").
		Order("id DESC").
		// Limit(500).
		Scan(&collaboratorWithSchedule).
		Error; err != nil {
		return nil, err
	}

	return collaboratorWithSchedule, nil
}

func GetCollaboratorForLeader(leaderDocument string) ([]entity.Collaborators, error) {
	collaboratorWithSchedule := []entity.Collaborators{}

	if err := config.DB.Table("collaborators").
		Select("*").
		Where("leader_document = ?", leaderDocument).
		Order("id DESC").
		// Limit(500).
		Scan(&collaboratorWithSchedule).
		Error; err != nil {
		return nil, err
	}

	return collaboratorWithSchedule, nil
}

func GetCollaborator(document string) (*models.Collaborators, error) {
	var collaborator models.Collaborators
	err := config.DB.Model(&models.Collaborators{}).Where("document = ?", document).First(&collaborator).Error
	if err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
		return nil, errors.New("Colaborador no encontrado")
	}
	return &collaborator, nil
}

func GetAllTranslatedService() ([]entity.Translatedcollaborators, error) {
	translatedcollaborators := []entity.Translatedcollaborators{}

	err := config.DB.
		Table("translatedcollaborators t").
		Select("t.*, c.f_name, c.l_name").
		Joins("INNER JOIN collaborators c ON t.fk_collaborator_id = c.id").
		Scan(&translatedcollaborators).Error
	if err != nil {
		return nil, err
	}

	return translatedcollaborators, nil
}

// func (s *CollaboratorService) GetByDocument(document string) (*models.Collaborators, error) {

// 	var collaborator models.Collaborators
// 	err := s.db.First(&collaborator, "document = ?", document).Error
// 	if err != nil {
// 		if errors.Is(err, gorm.ErrRecordNotFound) {
// 			return nil, fmt.Errorf("Colaborador no encontrado")
// 		}
// 		return nil, err
// 	}

// 	return &collaborator, nil

// }
