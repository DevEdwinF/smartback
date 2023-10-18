package user

import (
	"github.com/DevEdwinF/smartback.git/internal/app/controllers"
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/DevEdwinF/smartback.git/internal/config/middleware"
	"github.com/labstack/echo/v4"
)

func UserRoutes(e *echo.Echo) {

	userService := &services.UserService{}
	UserController := controllers.NewUserController(userService)

	group := e.Group("/api/user")
	group.POST("/create", UserController.CreateUser, middleware.AuthToken)
	group.GET("/all", UserController.GetAllUsers, middleware.AuthToken)
	group.GET("/:doc", UserController.GetUserById, middleware.AuthToken)
	group.PATCH("/update", UserController.UpdateUser, middleware.AuthToken)
	group.DELETE("/delete/:doc", UserController.DeleteUser, middleware.AuthToken)
}
