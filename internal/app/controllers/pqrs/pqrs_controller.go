package controllers

import (
	"net/http"

	services "github.com/DevEdwinF/smartback.git/internal/app/services/pqrs"
	"github.com/labstack/echo/v4"
)

func GetPqrsSacs(c echo.Context) error {
	pqrsSacs, err := services.GetPqrsSacs()
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return c.JSON(http.StatusOK, pqrsSacs)
}
