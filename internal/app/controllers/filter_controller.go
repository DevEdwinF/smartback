package controllers

import (
	"net/http"

	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/entity"
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

	filter := entity.CollaboratorFilter{}
	ctx.Bind(&filter)

	filter.SetDefault()

	collaborator, err := c.filterService.CollaboratorFilter(filter)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, map[string]interface{}{
			"error": "Colaborador no existe",
		})
	}

	return ctx.JSON(http.StatusOK, collaborator)
}
