package router

import "github.com/labstack/echo/v4"

func PqrsManageRoutes(e *echo.Echo) {

	group := e.Group("/api/pqrs")

	group.GET("/all", func(c echo.Context) error {
		return c.String(200, "Hola desde pqrs")
	})
}
