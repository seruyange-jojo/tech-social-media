package models

import "gorm.io/gorm"

type Message struct {
	gorm.Model
	SenderId uint `gorm:"not null"`
	ReceiverId uint `gorm:"not null"`
	Content string `gorm:"not null"`
}