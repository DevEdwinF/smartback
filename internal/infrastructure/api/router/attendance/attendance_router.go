package attendance

import (
	controllers "github.com/DevEdwinF/smartback.git/internal/app/controllers/attendance"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	group := e.Group("/api/attendance")
	group.POST("/register", controllers.SaveRegisterAttendance)
	group.GET("/all", controllers.GetAllAttendance)
}
