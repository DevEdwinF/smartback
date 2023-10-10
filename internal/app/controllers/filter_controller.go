package controllers

import (
	"net/http"

	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/labstack/echo/v4"
)

type Controller struct {
	filterService *services.FilterService
}

func NewController(filterService *services.FilterService) *Controller {
	return &Controller{
		filterService: filterService,
	}
}

func (c *Controller) CollaboratorFilterHandler(ctx echo.Context) error {

	firstName := ctx.QueryParam("firstName")
	lastName := ctx.QueryParam("lastName")
	email := ctx.QueryParam("email")
	state := ctx.QueryParam("state")
	leader := ctx.QueryParam("leader")
	subprocess := ctx.QueryParam("subprocess")
	headquarters := ctx.QueryParam("headquarters")
	position := ctx.QueryParam("position")

	collaborator, err := c.filterService.CollaboratorFilter(firstName, lastName, email, state, leader, subprocess, headquarters, position)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Colaborador no existe",
		})
	}

	return ctx.JSON(http.StatusOK, collaborator)
}
