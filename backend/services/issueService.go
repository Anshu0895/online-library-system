package services

import (
	"online-library-system/database"
	"online-library-system/models"
)

func CreateIssueRegistry(issueRegistry *models.IssueRegistry) error {
	return database.DB.Create(issueRegistry).Error
}

func GetIssueRegistries() ([]models.IssueRegistry, error) {
	var issueRegistries []models.IssueRegistry
	result := database.DB.Find(&issueRegistries)
	return issueRegistries, result.Error
}

func GetIssueRegistry(id uint) (*models.IssueRegistry, error) {
	var issueRegistry models.IssueRegistry
	result := database.DB.First(&issueRegistry, "issue_id = ?", id)
	return &issueRegistry, result.Error
}

func UpdateIssueStatus(issueRegistry *models.IssueRegistry) error {
	return database.DB.Save(issueRegistry).Error
}
