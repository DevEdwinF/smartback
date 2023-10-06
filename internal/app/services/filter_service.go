package services

import (
	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
)

type FilterService struct{}

func NewFilterService() *FilterService {
	return &FilterService{}
}

func (s *FilterService) CollaboratorFilter(firstName, lastName, email, state, leader, subprocess, headquarters, position string) (*models.Collaborators, error) {
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
}

/* func (s *FilterService) CollaboratorFilter(firstName, lastName, email, state, leader, subprocess, headquarters, position string) (*models.Collaborators, error) {
	collaborator := models.Collaborators{}
	query := config.DB.Table("collaborators").Select("*")

	if firstName != "" {
		query = query.Where("FName = ?", firstName)
	}

	if lastName != "" {
		query = query.Where("LName = ?", lastName)
	}

	if email != "" {
		query = query.Where("Email = ?", email)
	}

	if state != "" {
		query = query.Where("State = ?", state)
	}

	if leader != "" {
		query = query.Where("Leader = ?", leader)
	}

	if subprocess != "" {
		query = query.Where("Subprocess = ?", subprocess)
	}

	if headquarters != "" {
		query = query.Where("Headquarters = ?", headquarters)
	}

	if position != "" {
		query = query.Where("Position = ?", position)
	}

	err := query.Scan(&collaborator).Error
	if err != nil {
		return nil, err
	}

	return &collaborator, nil
}
*/
