package auth

import (
	controllers "github.com/DevEdwinF/smartback.git/internal/app/controllers/auth"
	"github.com/DevEdwinF/smartback.git/internal/config/middleware"
	"github.com/labstack/echo/v4"
)

func AuthRoutes(e *echo.Echo) {
	group := e.Group("/auth")
	group.POST("/login", controllers.Login)
	group.GET("/user-info", controllers.GetUserInfo, middleware.AuthToken) // Verifica que el middleware esté configurado aquí
}
