package http

import (
	"github.com/dsolyakin/task-tracker/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware(c *gin.Context) {
	header := c.GetHeader("Authorization")
	if len(header) < 8 {
		c.AbortWithStatusJSON(401, gin.H{"error": "Ошибка токена"})
		return
	}
	tokenString := header[7:]

	err := utils.ParseToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(401, gin.H{"error": "Ошибка токена"})
		return
	}
	c.Next()
}
