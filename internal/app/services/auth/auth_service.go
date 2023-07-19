package service

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func AuthenticateUser(email, password string) (*models.User, error) {
	var userModel models.User

	if err := config.DB.Table("users").Where("email = ? and password = ?", email, password).Scan(&userModel).Error; err != nil {
		return nil, err
	}

	if userModel.ID == 0 {
		return nil, echo.NewHTTPError(http.StatusNotFound, "Usuario o contraseña incorrecta")
	}

	return &userModel, nil
}

func GenerateToken(user *models.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = user.Email
	claims["name"] = user.Name
	claims["role"] = user.FkRoleId
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}

	fmt.Println("Generated Token:", tokenString) // Agregar este log para imprimir el token generado

	return tokenString, nil
}
