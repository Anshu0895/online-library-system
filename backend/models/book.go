package models

type BookInventory struct {
	ISBN            string `gorm:"primaryKey" json:"isbn"`
	LibID           uint   `json:"lib_id"`
	Title           string `json:"title"`
	Authors         string `json:"authors"`
	Publisher       string `json:"publisher"`
	Version         string `json:"version"`
	TotalCopies     uint   `json:"total_copies"`
	AvailableCopies uint   `json:"available_copies"`
}
