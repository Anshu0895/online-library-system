package models

type User struct {
	ID            uint   `gorm:"primaryKey" json:"id"`
	Name          string `json:"name"`
	Email         string `gorm:"unique" json:"email"`
	Password      string `json:"password"`
	ContactNumber string `json:"contact_number"`
	Role          string `json:"role"` // Role can be "LibraryOwner", "LibraryAdmin", "Reader"
	LibID         uint   `json:"lib_id"`
}
