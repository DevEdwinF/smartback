package attendance

import (
	"github.com/DevEdwinF/smartback.git/internal/app/controllers"
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/DevEdwinF/smartback.git/internal/config/middleware"
	"github.com/labstack/echo/v4"
)

func AttendanceRoutes(e *echo.Echo) {
	attendanceService := services.NewAttendanceService()
	attendanceController := controllers.NewAttendanceController(attendanceService)

	group := e.Group("/api/attendance")

	/* group.GET("/validate/:doc", controllers.ValidateCollaboratorController) */
	group.POST("/register", attendanceController.SaveRegisterAttendance)
	group.GET("/all", attendanceController.GetAllAttendance, middleware.AuthToken)
	group.GET("/late/all", attendanceController.GetAllAttendanceForLate, middleware.AuthToken)
	group.GET("/leader/all", attendanceController.GetAttendanceForLeader, middleware.AuthToken)
	group.GET("/leader/late/all", attendanceController.GetAttendanceForLeaderToLate, middleware.AuthToken)
	group.GET("/late/all", attendanceController.GetAllAttendanceForLate, middleware.AuthToken)
	group.POST("/register/translated", attendanceController.SaveTranslated)
	group.GET("/all/translated", controllers.GetAllTranslatedController, middleware.AuthToken)
}
