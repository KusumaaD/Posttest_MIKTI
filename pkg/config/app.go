package config

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ConnectDatabase connects to the SQLite database
func ConnectDatabase() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("bookstore.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }
    return db, nil
}