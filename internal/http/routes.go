package http

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitRoutes(r *gin.Engine, db *gorm.DB) {
	taskH := &TaskHandler{DB: db}
	catH := &CategoryHandler{DB: db}
	tagH := &TagHandler{DB: db}
	authH := &AuthHandler{DB: db}

	r.POST("/register", authH.CreateUserHandler)
	r.POST("/login", authH.LoginHandler)

	protected := r.Group("/")
	protected.Use(AuthMiddleware)

	protected.GET("/tasks", taskH.GetTaskListHandler)
	protected.POST("/tasks", taskH.CreateTaskHandler)
	protected.DELETE("/tasks/:id", taskH.DeleteTaskHandler)
	protected.GET("/tasks/:id", taskH.GetTaskIdHandler)
	protected.PUT("/tasks/:id", taskH.UpdateTaskHandler)

	protected.GET("/categories", catH.GetCategoryListHandler)
	protected.POST("/categories", catH.CreateCategoryHandler)

	protected.GET("/tags", tagH.GetTagListHandler)
	protected.POST("/tags", tagH.CreateTagHandler)

}
