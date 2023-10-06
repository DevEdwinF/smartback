package filter

import (
	"github.com/DevEdwinF/smartback.git/internal/app/controllers"
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/labstack/echo/v4"
)

func FilterRoutes(e *echo.Echo) {
	filterService := services.NewFilterService()
	filterController := controllers.NewController(filterService)

	group := e.Group("/api/filter")

	group.GET("/collaborator", filterController.CollaboratorFilterHandler)
}
