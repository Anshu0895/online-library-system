package models

type BookInventory struct {
	ISBN            string `gorm:"primaryKey"`
	LibID           uint
	Title           string
	Authors         string
	Publisher       string
	Version         string
	TotalCopies     uint
	AvailableCopies uint
}
