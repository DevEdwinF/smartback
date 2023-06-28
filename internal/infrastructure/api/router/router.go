package router

import (
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/attendance"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func GlobalRouter(e *echo.Echo) {
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		// AllowHeaders: []string{"*"},
	}))
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))
	// pqrs.PqrsManageRoutes(e)
	// auth.AuthRoutes(e)
	attendance.AuthRoutes(e)
}
