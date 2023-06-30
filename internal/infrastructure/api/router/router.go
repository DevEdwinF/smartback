package router

import (
	attendance "github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/attendance"
	collaborator "github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/collaborator"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/schedule"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/stats"
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
	attendance.AttendanceRoutes(e)
	collaborator.CollaboratorRoutes(e)
	schedule.ScheduleRouter(e)
	stats.StatsRoutes(e)

}
