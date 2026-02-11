package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	h := &TaskHandler{DB: db}

	r.GET("/tasks", h.GetTaskListHandler)

	r.POST("/tasks", h.CreateTaskHandler)

	r.DELETE("/tasks/:id", h.DeleteTaskHandler)

	r.GET("/tasks/:id", h.GetTaskIdHandler)

	r.PUT("/tasks/:id", h.UpdateTaskHandler)
}
