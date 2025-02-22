package routes

import (
	"tech-social-media/controllers"
	"tech-social-media/middleware"
	// "tech-social-media/middleware"
	"github.com/gin-gonic/gin"
)

func SetupRouter(router *gin.Engine) {
	api := router.Group("/api")
	{
		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.POST("/posts", controllers.CreatePost)
		// api.GET("/posts", controllers.FetchPosts)

		api.POST("/message", controllers.SendMessage)


		admin := api.Group("/admin")
		admin.Use(middleware.AdminAuthMiddleware())
		{
			admin.POST("/courses", controllers.CreateCourse)
		}

	}
}
