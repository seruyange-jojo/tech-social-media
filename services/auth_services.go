package services

import (
	"errors"
	"fmt"
	"tech-social-media/database"
	"tech-social-media/models"

	"golang.org/x/crypto/bcrypt"
)

func Register(user models.User) error {
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	user.Password = string(hashedPassword)

	if user.Role == "" {
		user.Role = "user"
	}

	result := database.DB.Create(&user)

	if result.Error != nil {
		return errors.New("user not created")
	}

	fmt.Println("User created successfully")

	return nil
}

func Login(user models.User) error {
	result := database.DB.Where("email = ?", user.Email).First(&user)

	if result.Error != nil {
		return errors.New("user not found")
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(user.Password), 8)

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(hashedPassword)); err != nil {
		return errors.New("invalid password")
	}

	return nil
}

