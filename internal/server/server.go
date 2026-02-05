package server

import (
	"github.com/dsolyakin/task-tracker/internal/handler"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	r.GET("/ping", handler.PingHandler)
	r.GET("/task", handler.GetTask)
	r.Run()
}
