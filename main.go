// main.go
package main

import (
	"log"
	"net/http"
	"os"
	"github.com/gin-gonic/gin"
	"github.com/anusornc/go-gorm/models"
	"github.com/anusornc/go-gorm/db"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dbType := os.Getenv("DB_TYPE")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	
	database, err := db.ConnectDatabase(dbType, dbUser, dbPassword, dbHost, dbPort ,dbName)
	if err != nil {
		log.Fatalf("failed to connect database: %v", err)
	}
	err = database.AutoMigrate(&models.Item{})
	if err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	itemRepo := models.NewItemRepository(database)

	r := gin.Default()

	r.GET("/items", itemRepo.GetItems)
	r.POST("/items", itemRepo.PostItem)

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"message": "Not found"})
	})

	// Run the server
	if err := r.Run(":5000"); err != nil {
		log.Fatalf("Server is not running: %v", err)
	}
}

