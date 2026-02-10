package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	r.GET("/tasks", func(c *gin.Context) {
		GetTaskListHandler(c, db)
	})

	r.POST("/tasks", func(c *gin.Context) {
		CreateTaskHandler(c, db)
	})

	r.DELETE("/tasks/:id", func(c *gin.Context) {
		DeleteTaskHandler(c, db)
	})

	r.GET("/tasks/:id", func(c *gin.Context) {
		GetTaskIdHandler(c, db)
	})

	r.PUT("/tasks/:id", func(c *gin.Context) {
		UpdateTaskHandler(c, db)
	})
}
