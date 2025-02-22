package controllers

import (
	"net/http"
	"path/filepath"
	"tech-social-media/models"
	"tech-social-media/services"

	"github.com/gin-gonic/gin"
)

func CreateCourse(c *gin.Context) {

	role := c.MustGet("role").(string)	
	if role != "admin" {
		c.JSON(401, gin.H{"error": "only admins can create courses"})
		return
	}

	var course models.Course
	c.BindJSON(&course)

	file, _ := c.FormFile("video")
	if  file == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "no video file Uploaded"})
		return
	}

	videoPath := filepath.Join("uploads", file.Filename)
	if err := c.SaveUploadedFile(file, videoPath); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error":"error saving video file"})
	}

	course.VideoUrl = videoPath


	err := services.CreateCourse(course)
	if err != nil {
		c.JSON(400, gin.H{"error": "course not created"})
		return
	}

	c.JSON(200, gin.H{"message": "course created successfully"})

}