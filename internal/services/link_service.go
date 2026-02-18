// Link business logic
package services

import (
	"github.com/vishwanththalla/linkme/internal/database"
	"github.com/vishwanththalla/linkme/internal/models"
)

func CreateLink(title, url string, userID uint) error {
	link := models.Link{
		Title:  title,
		URL:    url,
		UserID: userID,
	}

	return database.DB.Create(&link).Error
}

func UpdateLink(linkID uint, title, url string, userID uint) error {
	var link models.Link

	result := database.DB.Where("id = ? AND user_id = ?", linkID, userID).First(&link)
	if result.Error != nil {
		return result.Error
	}

	link.Title = title
	link.URL = url

	return database.DB.Save(&link).Error
}

func DeleteLink(linkID uint, userID uint) error {
	result := database.DB.Where("id = ? AND user_id = ?", linkID, userID).
		Delete(&models.Link{})

	return result.Error
}

func GetUserLinks(userID uint, page, limit int) ([]models.Link, error) {
	var links []models.Link

	offset := (page - 1) * limit

	err := database.DB.
		Where("user_id = ?", userID).
		Limit(limit).
		Offset(offset).
		Find(&links).Error

	return links, err
}
