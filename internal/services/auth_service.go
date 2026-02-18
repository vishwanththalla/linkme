// Auth business logic
package services

import (
	"errors"

	"github.com/vishwanththalla/linkme/internal/database"
	"github.com/vishwanththalla/linkme/internal/models"
	"github.com/vishwanththalla/linkme/internal/utils"
)

func RegisterUser(email, password string) error {
	var existingUser models.User

	result := database.DB.Where("email = ?", email).First(&existingUser)
	if result.RowsAffected > 0 {
		return errors.New("user already exists")
	}

	hashedPassword, err := utils.HashPassword(password)
	if err != nil {
		return err
	}

	user := models.User{
		Email:    email,
		Password: hashedPassword,
	}

	return database.DB.Create(&user).Error
}
