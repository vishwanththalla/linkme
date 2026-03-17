package database

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB // 🔥 THIS FIXES YOUR ERROR

func Connect() *gorm.DB {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using environment variables")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s",
		getEnv("DB_HOST", "db"),
		getEnv("DB_USER", "postgres"),
		getEnv("DB_PASSWORD", "postgres"),
		getEnv("DB_NAME", "linkme"),
		getEnv("DB_PORT", "5432"),
		getEnv("DB_SSLMODE", "disable"),
	)

	var errDB error

	for i := 0; i < 10; i++ {
		DB, errDB = gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if errDB == nil {
			log.Println("✅ Connected to database")
			return DB
		}

		log.Println("⏳ Waiting for database...")
		time.Sleep(3 * time.Second)
	}

	log.Fatal("❌ Failed to connect to database:", errDB)
	return nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
