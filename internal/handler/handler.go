package handler

import (
	"fmt"

	"github.com/dsolyakin/task-tracker/domain"
	"github.com/gin-gonic/gin"
)

func PingHandler(c *gin.Context) {
	fmt.Println("!!! Кто-то постучался на /ping !!!")
}

func GetTask(c *gin.Context) {
	myTask := domain.Task{
		ID:   1,
		Name: "test task",
	}
	c.JSON(200, myTask)
}
