package router

import (
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router/attendance"
	"github.com/labstack/echo/v4"
)

func GlobalRouter(e *echo.Echo) {
	// pqrs.PqrsManageRoutes(e)
	// auth.AuthRoutes(e)
	attendance.AuthRoutes(e)
}
