package stats

import (
	controllers "github.com/DevEdwinF/smartback.git/internal/app/controllers/stats"
	"github.com/DevEdwinF/smartback.git/internal/config/middleware"
	"github.com/labstack/echo/v4"
)

func StatsRoutes(e *echo.Echo) {

	group := e.Group("/api/stats")

	group.GET("/all", controllers.CountAttendancesAll, middleware.AuthToken)
	group.GET("/day/all", controllers.CountAttendanceDay, middleware.AuthToken)
}
