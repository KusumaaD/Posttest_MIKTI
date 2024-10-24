package main

import (
	"Posttest_MIKTI/pkg/config"
	"Posttest_MIKTI/pkg/models"
	"Posttest_MIKTI/pkg/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
    db, err := config.ConnectDatabase()
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // Migrate the database
    models.MigrateBooks(db)

    // Setup Router
    r := gin.Default()
    routes.BookStoreRoutes(r, db)

    // Run server
    r.Run(":8080")
}