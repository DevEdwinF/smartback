package router

import (
	controller "github.com/DevEdwinF/smartback.git/internal/app/controllers/pqrs"
	"github.com/labstack/echo/v4"
)

func PqrsManageRoutes(e *echo.Echo) {

	group := e.Group("/api/pqrs")

	group.GET("/all", controller.GetPqrsSacs)
}
