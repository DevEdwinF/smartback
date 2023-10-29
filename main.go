package main

import (
	"github.com/DevEdwinF/smartback.git/internal/app/services"
	"github.com/DevEdwinF/smartback.git/internal/config"
	"github.com/DevEdwinF/smartback.git/internal/infrastructure/api/router"
	"github.com/labstack/echo/v4"
)

func main() {
	go services.RunCronJob()

	config.ConnectDB()
	config.KactusDB()

	//changes
	e := echo.New()

	router.GlobalRouter(e)

	e.Logger.Fatal(e.Start(":8080"))
}
