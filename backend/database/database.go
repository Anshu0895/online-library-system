package database

import (
  "gorm.io/driver/sqlite"
  "gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
  database, err := gorm.Open(sqlite.Open("library.db"), &gorm.Config{})
  if err != nil {
    panic("Failed to connect to database")
  }

  database.AutoMigrate(&Library{}, &User{}, &BookInventory{}, &RequestEvent{}, &IssueRegistry{})
  DB = database
}
