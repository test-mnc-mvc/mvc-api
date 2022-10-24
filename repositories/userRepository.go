package repositories

import (
	"errors"
	_config "mvc/config"
	_entities "mvc/entities"
	_models "mvc/models"
)

func GetUsers() ([]_entities.User, error) {
	var users []_models.User
	result := _config.DB.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}
	responseData := []_entities.User{}
	for _, value := range users {
		var response _entities.User
		response.ID = value.ID
		response.Name = value.Name
		response.Email = value.Email
		response.CreatedAt = value.CreatedAt

		responseData = append(responseData, response)
	}
	return responseData, nil
}

func GetUserById(id int) (interface{}, error) {
	var users _models.User
	result := _config.DB.First(&users, id)

	responseData := _entities.UserResponseData{
		ID:        users.ID,
		Name:      users.Name,
		Email:     users.Email,
		CreatedAt: users.UpdatedAt,
	}
	if result.Error != nil {
		return nil, result.Error
	}
	return responseData, nil
}

func CreateUser(userReq _entities.UserRequestData) (int, error) {
	user := _models.User{
		Name:     userReq.Name,
		Email:    userReq.Email,
		Password: userReq.Password,
	}

	result := _config.DB.Create(&user)
	if result.Error != nil {
		return 0, result.Error
	}
	if result.RowsAffected == 0 {
		return 0, errors.New("failed to insert data")
	}
	return int(result.RowsAffected), nil
}
