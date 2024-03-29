package services

import (
	"errors"
	"math"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/DevEdwinF/smartback.git/internal/utils"
	"gorm.io/gorm"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) CreateUser(user entity.UserData) error {
	err := s.ValidateUser(user)
	if err != nil {
		return errors.New("El usuario ya existe")
	}

	err = config.DB.Table("users").Create(&user).Error
	if err != nil {
		return errors.New("failed to create user")
	}

	return nil
}

func (s *UserService) GetAllUsers(filter entity.UserFilter) (entity.Pagination, error) {

	offset := (filter.Page - 1) * filter.Limit
	var count int64
	var where string
	user := []models.User{}
	utils.BuildFilters("document", filter.Document, "OR", &where)
	utils.BuildFilters("f_name", filter.FName, "OR", &where)
	utils.BuildFilters("l_name", filter.LName, "OR", &where)
	utils.BuildFilters("email", filter.Email, "OR", &where)
	utils.BuildFilters("role_name", filter.RoleName, "OR", &where)

	err := config.DB.Table("users").
		Select("users.*, roles.name as role_name").
		Joins("INNER JOIN roles ON users.fk_role_id = roles.id").
		Where(where).
		Order("id DESC").
		Count(&count).
		Offset(offset).Limit(filter.Limit).
		Scan(&user).Error
	if err != nil {
		return entity.Pagination{}, err
	}

	return entity.Pagination{
		Page:      filter.Page,
		Limit:     filter.Limit,
		TotalPage: int(math.Ceil(float64(count) / float64(filter.Limit))),
		TotalRows: count,
		Rows:      user,
	}, nil

}

func (s *UserService) GetUserById(document string) (models.User, error) {
	var user models.User
	err := config.DB.Table("users").
		Select("users.*, roles.name as role_name").
		Joins("INNER JOIN roles ON users.fk_role_id = roles.id").
		Where("users.document = ?", document).
		First(&user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

// func (s *UserService) UpdateUser(user entity.UserData) error {
// 	err := config.DB.Table("users").Where("document = ?", user.Document).Updates(&user).Error
// 	if err != nil {
// 		return errors.New("failed to update user")
// 	}
// 	return nil
// }

func (s *UserService) UpdateUser(user entity.UserData) error {
	err := config.DB.Table("users").Where("document = ?", user.Document).Updates(user).Error
	if err != nil {
		return errors.New("failed to update user")
	}
	return nil
}

func (s *UserService) DeleteUser(document string) error {
	err := config.DB.Table("users").Where("document = ?", document).Delete(&models.User{}).Error
	if err != nil {
		return errors.New("failed to delete user")
	}
	return nil
}

func (s *UserService) ValidateUser(user entity.UserData) error {
	var existingUser entity.UserData
	err := config.DB.Table("users").
		Where("document = ?", user.Document).
		First(&existingUser).Error

	if err == gorm.ErrRecordNotFound {
		return nil
	} else if err != nil {
		return errors.New("database error")
	}

	return errors.New("El usuario ya existe")
}
