package services

import (
	"online-library-system/database"
	"online-library-system/models"
)

func CreateLibrary(library *models.Library) error {
	return database.DB.Create(library).Error
}

func GetLibraries() ([]models.Library, error) {
	var libraries []models.Library
	result := database.DB.Find(&libraries)
	return libraries, result.Error
}

func GetLibrary(id uint) (*models.Library, error) {
	var library models.Library
	result := database.DB.First(&library, "id = ?", id)
	return &library, result.Error
}

func UpdateLibrary(library *models.Library) error {
	return database.DB.Save(library).Error
}

func DeleteLibrary(id uint) error {
	var library models.Library
	result := database.DB.First(&library, "id = ?", id)
	if result.Error != nil {
		return result.Error
	}
	return database.DB.Delete(&library).Error
}
