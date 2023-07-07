package router

import (
	controller "github.com/DevEdwinF/smartback.git/internal/app/controllers/collaborator"
	"github.com/labstack/echo/v4"
)

func CollaboratorRoutes(e *echo.Echo) {

	group := e.Group("/api/collaborator")

	// group.POST("/save", controller.SaveCollaborator)
	group.GET("/all", controller.GetAllCollaborators)
	group.GET("/find/:document", controller.GetCollaborator)
	group.DELETE("/delete/:doc", controller.DeleteCollaborator)
}
