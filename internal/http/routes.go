package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	taskH := &TaskHandler{DB: db}
	catH := &CategoryHandler{DB: db}

	r.GET("/tasks", taskH.GetTaskListHandler)
	r.POST("/tasks", taskH.CreateTaskHandler)
	r.DELETE("/tasks/:id", taskH.DeleteTaskHandler)
	r.GET("/tasks/:id", taskH.GetTaskIdHandler)
	r.PUT("/tasks/:id", taskH.UpdateTaskHandler)

	r.GET("/categories", catH.GetCategoryListHandler)
	r.POST("/categories", catH.CreateCategoryHandler)
}
