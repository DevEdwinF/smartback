package router

import (
	controller "github.com/DevEdwinF/smartback.git/internal/app/controllers/collaborator"
	"github.com/DevEdwinF/smartback.git/internal/config/middleware"
	"github.com/labstack/echo/v4"
)

func CollaboratorRoutes(e *echo.Echo) {

	group := e.Group("/api/collaborator")

	// group.POST("/save", controller.SaveCollaborator)
	group.GET("/all", controller.GetAllCollaborators, middleware.AuthToken)
	group.GET("/find/:document", controller.GetCollaborator, middleware.AuthToken)
	group.DELETE("/delete/:doc", controller.DeleteCollaborator, middleware.AuthToken)
}
