// Link model
package models

import "gorm.io/gorm"

type Link struct {
	gorm.Model
	Title  string `gorm:"not null"`
	URL    string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
}

