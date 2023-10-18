package kactus

import (
	"github.com/DevEdwinF/smartback.git/internal/app/controllers"
	"github.com/DevEdwinF/smartback.git/internal/config/middleware"
	"github.com/labstack/echo/v4"
)

func KactusRouter(e *echo.Echo) {

	group := e.Group("/api/kactus")

	group.GET("/collaborators/all", controllers.GetAllColab, middleware.AuthToken)
	group.GET("/collaborators/:document", controllers.GetCollab, middleware.AuthToken)
}
