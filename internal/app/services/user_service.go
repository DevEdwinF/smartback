package services

import (
	"errors"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
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

func (s *UserService) GetAllUsers(users *[]models.User) error {
	err := config.DB.Table("users").
		Select("users.*, roles.name as role_name").
		Joins("INNER JOIN roles ON users.fk_role_id = roles.id").
		Find(users).Error
	if err != nil {
		return err
	}

	return nil
}

// func GetAllUsers() ([]models.User, error) {
// 	users := []models.User{}
// 	err := config.DB.Table("users").Joins("INNER JOIN roles ON users.fk_role_id = roles.id").Find(&users).Error
// 	if err != nil {
// 		return nil, err
// 	}

// 	return users, nil
// }

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
