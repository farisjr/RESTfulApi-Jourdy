package controllers

import (
	"farisjr/clean-code/lib/database"
	"farisjr/clean-code/models"
	"net/http"
	"strconv"

	"github.com/labstack/echo"
)

//Controller to get all users
func GetUserControllers(c echo.Context) error {
	users, err := database.GetUsers()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"status": "success",
		"users":  users,
	})
}

//Controller for get user by id
func GetUserByIdController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	user, err := database.GetUserById(id)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot fetch data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    user,
	})
}

//Controller for creating new user
func CreateUserController(c echo.Context) error {
	create_user := models.User{}
	c.Bind(&create_user)

	user, err := database.CreateUser(create_user)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success create new user",
		"user":    user,
	})
}

//Controller for deleting user by id
func DeleteUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}
	delete_user, err := database.GetUserById(id)
	c.Bind(&delete_user)
	userId, err := database.DeleteUserById(delete_user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot delete data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "delete success ",
		"data":    userId,
	})
}

//Controller for updating user data by id
func UpdateUserController(c echo.Context) error {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]interface{}{
			"message": "invalid id",
		})
	}

	update_user, err := database.GetUserById(id)
	c.Bind(&update_user)
	userId, err := database.UpdateUserById(update_user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]interface{}{
			"message": "cannot update data",
		})
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message": "success",
		"data":    userId,
	})
}
