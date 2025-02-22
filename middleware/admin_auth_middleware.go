package middleware

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		role := c.MustGet("role").(string)
		if role != "admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "only admins are allowed to perform this action"})
			c.Abort()
			return
		}
		c.Next()
	}
}