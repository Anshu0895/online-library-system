package services

import (
	"online-library-system/database"
	"online-library-system/models"
)

func CreateRequestEvent(requestEvent *models.RequestEvent) error {
	return database.DB.Create(requestEvent).Error
}

func GetRequestEvents() ([]models.RequestEvent, error) {
	var requestEvents []models.RequestEvent
	result := database.DB.Find(&requestEvents)
	return requestEvents, result.Error
}

func GetRequestEvent(id uint) (*models.RequestEvent, error) {
	var requestEvent models.RequestEvent
	result := database.DB.First(&requestEvent, "req_id = ?", id)
	return &requestEvent, result.Error
}

func ApproveRequestEvent(requestEvent *models.RequestEvent) error {
	requestEvent.ApprovalDate = "today's date" // Set to the current date in real implementation
	return database.DB.Save(requestEvent).Error
}
