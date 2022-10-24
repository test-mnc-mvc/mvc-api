package routes

import (
	_controllers "mvc/controllers"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/users", _controllers.GetUsersController)
	e.GET("/users/:id", _controllers.GetUserByIdController)
	e.POST("/users", _controllers.CreateUserController)

	return e
}
