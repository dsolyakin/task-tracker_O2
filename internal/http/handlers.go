package http

import (
	"fmt"

	"github.com/dsolyakin/task-tracker/domain"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

type TaskHandler struct {
	DB *gorm.DB
}

func (h *TaskHandler) CreateTaskHandler(c *gin.Context) {
	var task domain.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		fmt.Println("CreateTask. Ошибка парсинга json:", err)
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	result := h.DB.Create(&task)
	err = result.Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось сохранить в базу"})
		return
	}

	c.JSON(201, task)
}

func (h *TaskHandler) GetTaskListHandler(c *gin.Context) {
	var tasks []domain.Task

	result := h.DB.Find(&tasks)
	err := result.Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось получить список задач"})
		return
	}

	c.JSON(200, tasks)
}

func (h *TaskHandler) DeleteTaskHandler(c *gin.Context) {
	id := c.Param("id")

	result := h.DB.Delete(&domain.Task{}, id)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Задача не найдена"})
		return
	}
	c.Status(204)
}

func (h *TaskHandler) GetTaskIdHandler(c *gin.Context) {
	id := c.Param("id")

	var task domain.Task

	result := h.DB.First(&task, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Задача не найдена"})
		return
	}
	c.JSON(200, task)

}

func (h *TaskHandler) UpdateTaskHandler(c *gin.Context) {
	id := c.Param("id")

	var task domain.Task

	result := h.DB.First(&task, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Задача не найдена"})
		return
	}

	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	h.DB.Save(&task)

	c.JSON(200, task)
}
