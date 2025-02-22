package controllers

import (
	"tech-social-media/models"
	"tech-social-media/services"
	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context) {
	var message models.Message
	if  err := c.ShouldBindJSON(&message); err != nil {
		c.JSON(400, gin.H{"error" : "Invalid message data"})
		return
	}

	userId := c.MustGet("UserId").(uint)
	message.SenderId = userId

	err := services.SendMessage(&message)
	if err != nil {
		c.JSON(400, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(201, gin.H{"message" : "Message sent successfully"})

}