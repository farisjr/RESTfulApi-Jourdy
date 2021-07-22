package database

import (
	"farisjr/project/config"
	"farisjr/project/middlewares"
	"farisjr/project/models"
)

func GetUsers() (interface{}, error) {
	var users []models.User
	if err := config.DB.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func GetUserById(id int) (models.User, error) {
	var user models.User
	if err := config.DB.Find(&user, "id=?", id).Error; err != nil {
		return user, err
	}
	return user, nil
}

func CreateUser(create_user models.User) (interface{}, error) {
	if err := config.DB.Save(&create_user).Error; err != nil {
		return nil, err
	}
	return create_user, nil
}

func DeleteUserById(delete_user models.User) (interface{}, error) {
	if err := config.DB.Delete(&delete_user).Error; err != nil {
		return nil, err
	}
	return delete_user, nil
}

func UpdateUserById(update_user models.User) (interface{}, error) {

	if err := config.DB.Save(&update_user).Error; err != nil {
		return nil, err
	}
	return update_user, nil
}

func GetDetailUsers(userId int) (interface{}, error) {
	var user models.User
	if err := config.DB.Find(&user, userId).Error; err != nil {
		return nil, err
	}

	return user, nil
}

func LoginUsers(user *models.User) (interface{}, error) {
	var err error
	if err = config.DB.Where("email = ? AND password = ?", user.Email, user.Password).First(user).Error; err != nil {
		return nil, err
	}

	user.Token, err = middlewares.CreateToken(int(user.ID))
	if err != nil {
		return nil, err
	}

	if err := config.DB.Save(user).Error; err != nil {
		return nil, err
	}

	return user, nil
}
