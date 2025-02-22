package models

import "gorm.io/gorm"
type Post struct {
	gorm.Model
	UserID uint `gorm:"not null"`
	Content string `gorm:"not null"`
	// Likes int `gorm:"default:0"`
}