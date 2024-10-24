package controllers

import (
	"Posttest_MIKTI/pkg/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// GetBooks gets all books
func GetBooks(c *gin.Context, db *gorm.DB) {
    var books []models.Book
    db.Find(&books)
    c.JSON(http.StatusOK, books)
}

// CreateBook creates a new book
func CreateBook(c *gin.Context, db *gorm.DB) {
    var book models.Book
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Create(&book)
    c.JSON(http.StatusOK, book)
}

// GetBookByID retrieves a book by its ID
func GetBookByID(c *gin.Context, db *gorm.DB) {
    var book models.Book
    if err := db.First(&book, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    c.JSON(http.StatusOK, book)
}

// UpdateBook updates a book by its ID
func UpdateBook(c *gin.Context, db *gorm.DB) {
    var book models.Book
    if err := db.First(&book, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    if err := c.ShouldBindJSON(&book); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }
    db.Save(&book)
    c.JSON(http.StatusOK, book)
}

// DeleteBook deletes a book by its ID
func DeleteBook(c *gin.Context, db *gorm.DB) {
    var book models.Book
    if err := db.First(&book, c.Param("id")).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
        return
    }
    db.Delete(&book)
    c.JSON(http.StatusOK, gin.H{"message": "Book deleted"})
}