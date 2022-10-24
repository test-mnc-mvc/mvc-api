package controllers

import (
	"fmt"
	"net/http"
	"strconv"

	_entities "mvc/entities"
	_helper "mvc/helper"
	_repositories "mvc/repositories"

	"github.com/labstack/echo/v4"
)

func GetUsersController(c echo.Context) error {

	users, err := _repositories.GetUsers()

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get all data"))
	}

	responseData := []_entities.UserResponseData{}
	for _, value := range users {
		var response _entities.UserResponseData
		response.ID = value.ID
		response.Name = value.Name
		response.Email = value.Email
		response.CreatedAt = value.CreatedAt

		responseData = append(responseData, response)
	}
	return c.JSON(http.StatusOK, _helper.ResponseSuccessWithData("success to get all data", responseData))
}

func GetUserByIdController(c echo.Context) error {
	id := c.Param("id")
	iduser, errid := strconv.Atoi(id)
	if errid != nil {
		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed to get data, id not recognize"))
	}
	users, err := _repositories.GetUserById(iduser)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to get user data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseSuccessWithData("success to get data", users))
}

func CreateUserController(c echo.Context) error {
	user := _entities.UserRequestData{}
	err := c.Bind(&user)
	if err != nil {
		fmt.Println("error", err)

		return c.JSON(http.StatusBadRequest, _helper.ResponseFailed("failed create user bind"))
	}

	result, errCreate := _repositories.CreateUser(user)

	if errCreate != nil {

		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert data"))
	}
	if result == 0 {
		return c.JSON(http.StatusInternalServerError, _helper.ResponseFailed("failed to insert data"))
	}

	return c.JSON(http.StatusOK, _helper.ResponseSuccessNoData("success insert data"))

}
