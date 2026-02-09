package http

import (
	"github.com/dsolyakin/task-tracker/domain"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func BuHandler(c *gin.Context, db *gorm.DB) {
	var task domain.Task

	result := db.First(&task)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "В базе пока нет ни одной задачи"})
		return

	}
	c.JSON(200, task)
}
