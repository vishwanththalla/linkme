package main

import (
	"log"

	"github.com/joho/godotenv"
	"github.com/vishwanththalla/linkme/internal/database"
	"github.com/vishwanththalla/linkme/internal/models"
	"github.com/vishwanththalla/linkme/internal/utils"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	database.Connect()

	log.Println("Seeding database...")

	// Clear existing data
	database.DB.Exec("TRUNCATE TABLE links RESTART IDENTITY CASCADE")
	database.DB.Exec("TRUNCATE TABLE users RESTART IDENTITY CASCADE")

	// Create users
	password1, _ := utils.HashPassword("password123")
	password2, _ := utils.HashPassword("password123")

	user1 := models.User{
		Email:    "test1@example.com",
		Password: password1,
	}

	user2 := models.User{
		Email:    "test2@example.com",
		Password: password2,
	}

	database.DB.Create(&user1)
	database.DB.Create(&user2)

	// Create links
	links := []models.Link{
		{Title: "Google", URL: "https://google.com", UserID: user1.ID},
		{Title: "GitHub", URL: "https://github.com", UserID: user1.ID},
		{Title: "YouTube", URL: "https://youtube.com", UserID: user2.ID},
	}

	for _, link := range links {
		database.DB.Create(&link)
	}

	log.Println("Seeding complete.")
}
