package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context){
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error" : "unauthorized acess"})
			c.Abort()
			return
		}
		c.Next()
	}
}

