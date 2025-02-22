package services

import (
	"errors"
	"tech-social-media/models"
	"tech-social-media/database"
)

func CreateCourse(course models.Course) error {
	result := database.DB.Create(&course)

	if result.Error != nil {
		return errors.New("could not create course")
	}

	return nil
}