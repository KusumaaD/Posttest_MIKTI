package routes

import (
	"Posttest_MIKTI/pkg/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// BookStoreRoutes sets up the routes for the bookstore
func BookStoreRoutes(router *gin.Engine, db *gorm.DB) {
    router.GET("/book/", func(c *gin.Context) { controllers.GetBooks(c, db) })
    router.POST("/book/", func(c *gin.Context) { controllers.CreateBook(c, db) })
    router.GET("/book/:id", func(c *gin.Context) { controllers.GetBookByID(c, db) })
    router.PUT("/book/:id", func(c *gin.Context) { controllers.UpdateBook(c, db) })
    router.DELETE("/book/:id", func(c *gin.Context) { controllers.DeleteBook(c, db) })
}