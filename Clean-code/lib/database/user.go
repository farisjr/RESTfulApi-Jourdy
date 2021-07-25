package database

import (
	"farisjr/clean-code/config"
	"farisjr/clean-code/models"
)

//function get all user table
func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

//function to get user from id
func GetUserById(id int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

//function for creating new user
func CreateUser(create_user models.User) (interface{}, error) {
	if err := config.DB.Save(&create_user).Error; err != nil {
		return nil, err
	}
	return create_user, nil
}

//function for deleting user by id
func DeleteUserById(delete_user models.User) (interface{}, error) {
	if err := config.DB.Delete(&delete_user).Error; err != nil {
		return nil, err
	}
	return delete_user, nil
}

//function update user data by id
func UpdateUserById(update_user models.User) (interface{}, error) {

	if err := config.DB.Save(&update_user).Error; err != nil {
		return nil, err
	}
	return update_user, nil
}
