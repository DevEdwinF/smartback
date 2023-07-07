package schedule

import (
	controllers "github.com/DevEdwinF/smartback.git/internal/app/controllers/schedule"
	"github.com/labstack/echo/v4"
)

func ScheduleRouter(e *echo.Echo) {

	group := e.Group("/api/schedule")

	// group.POST("/save", controllers.SaveSchedule)
	group.GET("/all", controllers.GetAllCollaboratorsSchedule)
	group.DELETE("/delete/:id", controllers.DeleteSchedule)
	group.POST("/add", controllers.AssignScheduleToCollaborator)

}
