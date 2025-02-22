package services

import (
	"errors"
	"tech-social-media/models"
	"tech-social-media/database"
)

func SendMessage(message *models.Message) error{

	result := database.DB.Create(&message)

	if result.Error != nil{
		return errors.New("failed to send message")
	}

	return nil
}