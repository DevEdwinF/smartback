package attendance

import (
	"github.com/DevEdwinF/smartback.git/internal/app/controllers"
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/labstack/echo/v4"
)

func AttendanceRoutes(e *echo.Echo) {
	attendanceService := services.NewAttendanceService()
	attendanceController := controllers.NewAttendanceController(attendanceService)

	group := e.Group("/api/attendance")
	group.GET("/validate/:doc", controllers.ValidateCollaboratorController)
	group.POST("/register", attendanceController.SaveRegisterAttendance)
	// group.POST("/register", controllers.SaveRegisterAttendance)
	group.GET("/all", attendanceController.GetAllAttendance)
	group.POST("/register/translated", controllers.SaveTranslated)
	group.GET("/all/translated", controllers.GetAllTranslatedController)
}
