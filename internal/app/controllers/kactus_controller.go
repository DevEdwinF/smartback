package controllers

import (
	"net/http"

	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/labstack/echo/v4"
)

type BiEmple struct {
	CodEmpl string
}

func GetTest(c echo.Context) error {
	colaborador := []BiEmple{}

	config.KDB.Table("bi_emple").Select("*").Scan(&colaborador)
	return c.JSON(http.StatusOK, colaborador)
}
