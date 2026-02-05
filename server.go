package tasktracker

import (
	"github.com/dsolyakin/task-tracker/pkg/handler"
	"github.com/gin-gonic/gin"
)

func StartServer() {
	r := gin.Default()
	r.GET("/ping", handler.PingHandler)
	r.Run()
}
