package controllers

import (
	"net/http"

	service "github.com/DevEdwinF/smartback.git/internal/app/services/auth"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func Login(c echo.Context) error {
	userEntity := entity.User{}
	err := c.Bind(&userEntity)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	userModel, err := service.AuthenticateUser(userEntity.Email, userEntity.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	token, err := service.GenerateToken(userModel)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Error generating token")
	}

	return c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
		// "user":  userModel,
	})
}

func GetUserInfo(c echo.Context) error {
	user := c.Get("userToken").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	return c.JSON(http.StatusOK, map[string]interface{}{
		"email": claims["email"],
		"name":  claims["name"],
		"role":  claims["role"],
	})
}

func ValidateToken(c echo.Context) error {
	user := c.Get("userToken").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)

	rol, ok := claims["role"]
	if !ok {
		return echo.NewHTTPError(http.StatusBadRequest, "Error getting role")
	}
	return c.JSON(http.StatusOK, rol)
}
