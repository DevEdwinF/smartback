package services

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

	fmt.Println("Generated Token:", tokenString)

	return tokenString, nil
}

var blackListTokens = []string{}

func InvalidateToken(token string) {
	blackListTokens = append(blackListTokens, token)
}

func Logout(c echo.Context) error {
	authHeader := c.Request().Header.Get("authorization")

	if authHeader == "" || authHeader[:6] != "Bearer" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	tokenString := authHeader[7:]
	InvalidateToken(tokenString)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Sesión cerrada con exito",
	})
}

type Service struct {
	blackListTokens []string
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) invalidateToken(token string) {
	s.blackListTokens = append(s.blackListTokens, token)
}

func (s *Service) Logout(c echo.Context) error {
	authHeader := c.Request().Header.Get("Authorization")

	if authHeader == "" || len(authHeader) < 7 || authHeader[:7] != "Bearer " {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	tokenString := authHeader[7:]
	s.invalidateToken(tokenString)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Sesión cerrada con éxito",
	})
}