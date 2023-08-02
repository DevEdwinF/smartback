package attendance

import (
	"github.com/DevEdwinF/smartback.git/internal/app/controllers"
	"github.com/labstack/echo/v4"
)

func AttendanceRoutes(e *echo.Echo) {
	group := e.Group("/api/attendance")
	group.POST("/register", controllers.SaveRegisterAttendance)
	group.GET("/validate/:doc", controllers.ValidateCollaboratorController)
	// group.POST("/register/translated", controllers.SaveTranslated)
	group.GET("/all", controllers.GetAllAttendance)
}
