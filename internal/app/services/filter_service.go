package services

import (
	"math"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/DevEdwinF/smartback.git/internal/utils"
)

type FilterService struct{}

func NewFilterService() *FilterService {
	return &FilterService{}
}

/* func (s *FilterService) CollaboratorFilter(firstName, lastName, email, state, leader, subprocess, headquarters, position string) (*models.Collaborators, error) {
	collaborator := models.Collaborators{}
	query := config.DB.Model(&models.Collaborators{})

	switch {
	case firstName != "":
		query = query.Where("f_name ILIKE ?", "%"+firstName+"%")
	case lastName != "":
		query = query.Where("l_name ILIKE ?", "%"+lastName+"%")
	case email != "":
		query = query.Where("email ILIKE ?", "%"+email+"%")
	case state != "":
		query = query.Where("state ILIKE ?", "%"+state+"%")
	case leader != "":
		query = query.Where("leader ILIKE ?", "%"+leader+"%")
	case subprocess != "":
		query = query.Where("subprocess ILIKE ?", "%"+subprocess+"%")
	case headquarters != "":
		query = query.Where("headquarters ILIKE ?", "%"+headquarters+"%")
	case position != "":
		query = query.Where("position ILIKE ?", "%"+position+"%")
	}

	err := query.First(&collaborator).Error
	if err != nil {
		return nil, err
	}

	return &collaborator, nil
} */

func (s *FilterService) CollaboratorFilter(filter entity.CollaboratorFilter) (entity.Pagination, error) {
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
