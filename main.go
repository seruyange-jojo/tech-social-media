package main

import (

	"tech-social-media/config"
	"tech-social-media/models"
	"tech-social-media/routes"

	"github.com/gin-gonic/gin"
)

func main(){
	config.ConnectDB()
	config.DB.AutoMigrate(&models.User{}, &models.Post{}, &models.Comment{}, &models.Message{}, &models.Course{})

	server := gin.Default()

	routes.SetupRouter(server)

	// server.GET("/", func(c *gin.Context){
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message": "This application is moving on smoothly",
	// 	})
	// })

	server.Run(":8080")

}