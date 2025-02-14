package services

import (
	"online-library-system/database"
	"online-library-system/models"
)

func GetUsers() ([]models.User, error) {
	var users []models.User
	result := database.DB.Find(&users)
	return users, result.Error
}

func GetUser(id uint) (*models.User, error) {
	var user models.User
	result := database.DB.First(&user, "id = ?", id)
	return &user, result.Error
}

func UpdateUser(user *models.User) error {
	return database.DB.Save(user).Error
}

func DeleteUser(id uint) error {
	var user models.User
	result := database.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return database.DB.Delete(&user).Error
}
