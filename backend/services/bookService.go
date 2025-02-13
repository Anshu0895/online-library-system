package services

import (
	"online-library-system/database"
	"online-library-system/models"
)

func AddBook(book *models.BookInventory) error {
	return database.DB.Create(book).Error
}

func GetBooks() ([]models.BookInventory, error) {
	var books []models.BookInventory
	result := database.DB.Find(&books)
	return books, result.Error
}

func GetBook(isbn string) (*models.BookInventory, error) {
	var book models.BookInventory
	result := database.DB.First(&book, "isbn = ?", isbn)
	return &book, result.Error
}

func UpdateBook(book *models.BookInventory) error {
	return database.DB.Save(book).Error
}

func RemoveBook(isbn string) error {
	var book models.BookInventory
	result := database.DB.First(&book, "isbn = ?", isbn)
	if result.Error != nil {
		return result.Error
	}
	return database.DB.Delete(&book).Error
}
