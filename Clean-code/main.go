package main

import (
	"farisjr/clean-code/config"
	"farisjr/clean-code/middlewares"
	"farisjr/clean-code/routes"
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()   // create variable then assign with echo pkg
	config.InitDb()   // Doing database initialization
	config.InitPort() // Doing port initialization
	middlewares.LogMiddlewares(e)
	routes.New(e)
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", config.HTTP_PORT)))
}
