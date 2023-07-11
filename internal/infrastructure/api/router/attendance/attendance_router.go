package attendance

import (
	controllers "github.com/DevEdwinF/smartback.git/internal/app/controllers/attendance"
	"github.com/labstack/echo/v4"
)

func AttendanceRoutes(e *echo.Echo) {
	group := e.Group("/api/attendance")
	group.POST("/register", controllers.SaveRegisterAttendance)
	group.GET("/validate/:doc", controllers.ValidateColaborator)
	group.POST("/register/translated", controllers.SaveTranslated)
	group.GET("/all", controllers.GetAllAttendance)
}
