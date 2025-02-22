package database

import (
	"fmt"
	"os"
	"log"
	"tech-social-media/config"
	"tech-social-media/models"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
	

)

var DB *gorm.DB

func ConnectDB() {
	
	config.LoadEnv()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
	os.Getenv("DB_HOST"),
	os.Getenv("DB_USER"),
	os.Getenv("DB_PASSWORD"),
	os.Getenv("DB_NAME"),
	os.Getenv("DB_PORT"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database ", err)
	}

	log.Println("Database connected successfully")

	DB = db
	DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Message{}, &models.Course{})
}

