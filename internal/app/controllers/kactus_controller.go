package controllers

import (
	"net/http"

	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/labstack/echo/v4"
)

func GetAllColab(c echo.Context) error {
	collaborators, err := services.GetAllColab()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, collaborators)
}
