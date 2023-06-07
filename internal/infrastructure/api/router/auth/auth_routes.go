package auth

import (
	controllers "github.com/DevEdwinF/smartback.git/internal/app/controllers/auth"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	e.POST("/login", controllers.AuthController)
}
