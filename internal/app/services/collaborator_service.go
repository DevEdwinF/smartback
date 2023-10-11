package services

import (
	"errors"
	"fmt"
	"math"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/DevEdwinF/smartback.git/internal/utils"
	"gorm.io/gorm"
)

type CollaboratorService struct {
	db *gorm.DB
}

func NewCollaboratorService(db *gorm.DB) *CollaboratorService {
	return &CollaboratorService{db: db}
}

func ValidateCollaborator() ([]entity.Collaborators, error) {
	collaboratorWithSchedule := []entity.Collaborators{}

	if err := config.DB.Table("collaborators").
		Select("document").
		Order("id DESC").
		Scan(&collaboratorWithSchedule).
		Error; err != nil {
		return nil, err
	}

	return collaboratorWithSchedule, nil
}

func GetCollaboratorPage(filter entity.CollaboratorFilter) (entity.Pagination, error) {
	offset := (filter.Page - 1) * filter.Limit
	var count int64

	collaborator := []models.Collaborators{}
	var where string
	utils.BuildFilters("document", filter.Document, "OR", &where)
	utils.BuildFilters("f_name", filter.FName, "OR", &where)
	utils.BuildFilters("l_name", filter.LName, "OR", &where)
	utils.BuildFilters("bmail", filter.Bmail, "OR", &where)
	utils.BuildFilters("email", filter.Email, "OR", &where)
	utils.BuildFilters("position", filter.Position, "OR", &where)
	utils.BuildFilters("state", filter.State, "OR", &where)
	utils.BuildFilters("leader", filter.Leader, "OR", &where)
	utils.BuildFilters("leader_document", filter.LeaderDocument, "OR", &where)
	utils.BuildFilters("subprocess", filter.Subprocess, "OR", &where)
	utils.BuildFilters("headquarters", filter.Headquarters, "OR", &where)

	err := config.DB.Table("collaborators").Select("*").
		Where(where).
		Order("id DESC").
		Count(&count).
		Offset(offset).Limit(filter.Limit).
		Scan(&collaborator).Error
	if err != nil {
		return entity.Pagination{}, err
	}

	return entity.Pagination{
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalPage: int(math.Ceil(float64(count) / float64(filter.Limit))),
		TotalRows: count,
		Rows:      collaborator,
	}, nil
}

func GetCollaboratorForLeader(filter entity.CollaboratorFilter) (entity.Pagination, error) {
	offset := (filter.Page - 1) * filter.Limit
	var count int64

	collaborator := []models.Collaborators{}

	var where string
	utils.BuildFilters("document", filter.Document, "OR", &where)
	utils.BuildFilters("f_name", filter.FName, "OR", &where)
	utils.BuildFilters("l_name", filter.LName, "OR", &where)
	utils.BuildFilters("bmail", filter.Bmail, "OR", &where)
	utils.BuildFilters("email", filter.Email, "OR", &where)
	utils.BuildFilters("position", filter.Position, "OR", &where)
	utils.BuildFilters("state", filter.State, "OR", &where)
	utils.BuildFilters("leader", filter.Leader, "OR", &where)
	utils.BuildFilters("subprocess", filter.Subprocess, "OR", &where)
	utils.BuildFilters("headquarters", filter.Headquarters, "OR", &where)
	utils.BuildFilters("leader_document", filter.LeaderDocument, "AND", &where)

	err := config.DB.Table("collaborators").
		Select("*").
		Table("collaborators").
		Where(where).
		Order("id DESC").
		Count(&count).
		Offset(offset).Limit(filter.Limit).
		Scan(&collaborator).Error

	fmt.Print("esta es la cuenta: ", count)

	if err != nil {
		return entity.Pagination{}, err
	}

	return entity.Pagination{
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalPage: int(math.Ceil(float64(count) / float64(filter.Limit))),
		TotalRows: count,
		Rows:      collaborator,
	}, nil
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
