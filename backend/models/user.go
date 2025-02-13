package models

type User struct {
	ID            uint `gorm:"primaryKey"`
	Name          string
	Email         string `gorm:"unique"`
	ContactNumber string
	Role          string // Role can be "LibraryOwner", "LibraryAdmin", "Reader"
	LibID         uint
}
