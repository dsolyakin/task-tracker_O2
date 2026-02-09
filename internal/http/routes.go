package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/bu", func(c *gin.Context) {
		BuHandler(c, db)
	})
}
