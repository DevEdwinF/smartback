package services

import (
	"net/http"
	"time"

	"github.com/DevEdwinF/smartback.git/internal/app/models"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Auth(c echo.Context) error {
	userEntity := entity.User{}

	err := c.Bind(&userEntity)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	var userModel models.User

	if err := config.DB.Table("users").Where("email = ? and pass = ?", userEntity.Email, userEntity.Pass).Scan(&userModel).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	if userModel.ID == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Usuario o contraseña incorrectos")
	}

	if userModel.Pass != userEntity.Pass {
		return echo.NewHTTPError(http.StatusNotFound, "Usuario o contraseña incorrectos")
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = userModel.Email
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenSting, _ := token.SignedString([]byte("secret"))

	return c.JSON(http.StatusOK, map[string]string{
		"token": tokenSting,
	})

}

var blackListTokens = []string{}

func invalidateToken(token string) {
	blackListTokens = append(blackListTokens, token)
}

func Logout(c echo.Context) error {
	authHeader := c.Request().Header.Get("authorization")

	if authHeader == "" || authHeader[:6] != "Bearer" {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid token")
	}

	tokenString := authHeader[7:]
	invalidateToken(tokenString)
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Sesión cerrada con exito",
	})
}
