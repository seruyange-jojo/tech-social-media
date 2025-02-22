package controllers

import (
	"tech-social-media/models"
	"tech-social-media/services"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context) {	

	var user models.User
	c.BindJSON(&user)
	
	err := services.Register(user)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	
	c.JSON(200, gin.H{"message": "User registered successfully"})
}

func Login(c *gin.Context){
	var user models.User
	c.BindJSON(&user)

	err := services.Login(user)
	if err != nil {
		c.JSON(400, gin.H{"error": "Invalid login input"})
	}

	c.JSON(200, gin.H{"message": "User Logged in successfully"})
}