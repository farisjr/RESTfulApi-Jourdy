package routes

import (
	"farisjr/project/constants"
	"farisjr/project/controllers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func New(e *echo.Echo) {

	e.POST("/users", controllers.CreateUserController)
	e.POST("/login", controllers.LoginUsersController)
	e.GET("/books", controllers.GetBookControllers)
	e.GET("/books/:id", controllers.GetBookByIdController)

	//JWT Group
	r := e.Group("/jwt")
	r.Use(middleware.JWT([]byte(constants.SECRET_JWT)))

	r.GET("/users/:id", controllers.GetUserDetailController)
	r.GET("/users", controllers.GetUserControllers)
	r.PUT("/users/:id", controllers.UpdateUserController)
	r.DELETE("/users/:id", controllers.DeleteUserController)

	r.POST("/books", controllers.CreateBookController)
	r.DELETE("/books/:id", controllers.DeleteBookController)
	r.PUT("/books/:id", controllers.UpdateBookController)

	//return e
	// 	eAuth := e.Group("")
	// 	eAuth.Use(middleware.BasicAuth(middlewares.BasicAuthDb))
	// 	eAuth.GET("/rahasia", func(c echo.Context) error {
	// 		return c.JSON(http.StatusOK, map[string]interface{}{
	// 			"message": "rahasia, cuma bisa diakses yang login",
	// 		})
	// 	})
}
