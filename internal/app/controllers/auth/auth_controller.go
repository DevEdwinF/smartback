package controllers

import (
	"fmt"
	"net/http"

	services "github.com/DevEdwinF/smartback.git/internal/app/services/auth"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
	"github.com/labstack/echo/v4"
)

func AuthController(c echo.Context) error {
	userEntity := entity.User{}
	if err := c.Bind(&userEntity); err != nil {
		return err
	}

	token, err := services.AuthService(&userEntity)
	if err != nil {
		if err.Error() == "Usuario o contraseña incorrecta" {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return err
	}

	fmt.Println(token)
	return c.JSON(http.StatusOK, map[string]string{
		"token": token,
	})
}
