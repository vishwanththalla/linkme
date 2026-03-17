package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/vishwanththalla/linkme/internal/database"
	"github.com/vishwanththalla/linkme/internal/handlers"
	"github.com/vishwanththalla/linkme/internal/middleware"
	"github.com/vishwanththalla/linkme/internal/models"
)

func main() {

	// Load .env (optional inside Docker)
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	// Connect DB
	db := database.Connect()

	// Auto migrate
	err = db.AutoMigrate(
		&models.User{},
		&models.Link{},
	)
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	router := gin.Default()

	// CORS CONFIG
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000", "http://localhost"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Server running 🚀"})
	})

	router.POST("/register", handlers.Register)
	router.POST("/login", handlers.Login)

	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/links", handlers.CreateLink)
		auth.GET("/links", handlers.GetLinks)
		auth.PUT("/links/:id", handlers.UpdateLink)
		auth.DELETE("/links/:id", handlers.DeleteLink)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	router.Run(":" + port)
}
