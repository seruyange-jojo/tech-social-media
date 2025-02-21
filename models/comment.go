package models

import "gorm.io/gorm"

type Comment struct {
	gorm.Model
	PostId uint `gorm:"not null"`
	UserId uint `gorm:"not null"`
	Content string `gorm:"not null"`
}
