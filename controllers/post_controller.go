package controllers

import (
	"net/http"
	"tech-social-media/models"
	"tech-social-media/services"

	"github.com/gin-gonic/gin"
)

func CreatePost(c *gin.Context) {
	var post models.Post
	if  err := c.ShouldBindJSON(&post); err != nil { 
		c.JSON(http.StatusBadRequest, gin.H{"error" : "Invalid Post data"})
		return
	}

	userId := c.MustGet("UserId").(uint)
	post.UserID = userId

	err := services.CreatePost(post)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message" : "Post created successfully"})

}