package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/vishwanththalla/linkme/internal/database"
    "github.com/vishwanththalla/linkme/internal/handlers"
    "github.com/vishwanththalla/linkme/internal/middleware"


)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	database.Connect()

	router := gin.Default()
    
    router.POST("/login", handlers.Login)

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server running 🚀"})
	})

    router.POST("/register", handlers.Register)


    auth := router.Group("/")
    auth.Use(middleware.AuthMiddleware())
    {
        auth.POST("/links", handlers.CreateLink)
        auth.GET("/links", handlers.GetLinks)
        auth.PUT("/links/:id", handlers.UpdateLink)
        auth.DELETE("/links/:id", handlers.DeleteLink)
    }
		port = "8080"
	}

	router.Run(":" + port)
}
