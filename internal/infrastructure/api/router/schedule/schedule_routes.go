package schedule

import (
	controllers "github.com/DevEdwinF/smartback.git/internal/app/controllers/schedule"
	"github.com/DevEdwinF/smartback.git/internal/config/middleware"
	"github.com/labstack/echo/v4"
)

func ScheduleRouter(e *echo.Echo) {

	group := e.Group("/api/schedule")

	// group.POST("/save", controllers.SaveSchedule)
	group.GET("/all", controllers.GetAllCollaboratorsSchedule, middleware.AuthToken)
	group.DELETE("/delete/:id", controllers.DeleteSchedule, middleware.AuthToken)
	group.POST("/assign", controllers.AssignSchedulesToCollaborator, middleware.AuthToken)

}
