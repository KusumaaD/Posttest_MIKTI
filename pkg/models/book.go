package models

import "gorm.io/gorm"

// Book model represents a book in the bookstore
type Book struct {
    ID     uint    `gorm:"primaryKey"`
    Title  string  `json:"title"`
    Author string  `json:"author"`
    Price  float64 `json:"price"`
}

// MigrateBooks runs the migration for the book model
func MigrateBooks(db *gorm.DB) {
    db.AutoMigrate(&Book{})
}