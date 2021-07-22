package routes

import (
	"farisjr/project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUsersController)
	e.GET("/books", controllers.GetBookControllers)
	e.GET("/books/:id", controllers.GetBookByIdController)

	e.GET("/users/:id", controllers.GetUserDetailController)
	e.GET("/users", controllers.GetUserControllers)
	e.PUT("/users/:id", controllers.UpdateUserController)
	e.DELETE("/users/:id", controllers.DeleteUserController)

	e.POST("/books", controllers.CreateBookController)
	e.DELETE("/books/:id", controllers.DeleteBookController)
	e.PUT("/books/:id", controllers.UpdateBookController)


}
