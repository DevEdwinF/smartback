package services

import (
	"errors"
	"time"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/golang-jwt/jwt"
)

func AuthService(userEntity *entity.User) (string, error) {
	userModel := models.User{}
	if err := config.DB.Table("users").Where("email = ?", userEntity.Email).Scan(&userModel).Error; err != nil {
		return "", err
	}

	if userModel.ID == 0 {
		return "", errors.New("Usuario o contraseña incorrectos")
	}

	if userModel.Pass != userEntity.Pass {
		return "", errors.New("Usuario o contraseña incorrectos")
	}

	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = userModel.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenString, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
