package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/vishwanththalla/linkme/internal/database"
    "github.com/vishwanththalla/linkme/internal/handlers"

)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	router := gin.Default()
    

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server running 🚀"})
	})
    
    router.POST("/register", handlers.Register)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
