package services

import (
	"errors"
	"tech-social-media/models"
	"tech-social-media/database"
)

func CreatePost(post models.Post) error{

	result := database.DB.Create(&post)

	if result.Error != nil{
		return errors.New("failed to create post")
	}

	return nil
}