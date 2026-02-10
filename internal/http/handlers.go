package http

import (
	"fmt"

	"github.com/dsolyakin/task-tracker/domain"
	"gorm.io/gorm"

	"github.com/gin-gonic/gin"
)

func CreateTaskHandler(c *gin.Context, db *gorm.DB) {
	var task domain.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		fmt.Println("Ошибка парсинга JSON:", err)
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	err = db.Create(&task).Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось сохранить в базу"})
		return
	}

	c.JSON(201, task)
}

func GetTaskListHandler(c *gin.Context, db *gorm.DB) {
	var tasks []domain.Task

	result := db.Find(&tasks)
	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Не удалось получить список задач"})
		return
	}

	c.JSON(200, tasks)
}

func DeleteTaskHandler(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	result := db.Delete(&domain.Task{}, id)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Задача не найдена"})
		return
	}
	c.Status(204)
}

func GetTaskIdHandler(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	var task domain.Task

	result := db.First(&task, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Задача не найдена"})
		return
	}
	c.JSON(200, task)

}

func UpdateTaskHandler(c *gin.Context, db *gorm.DB) {
	id := c.Param("id")

	var task domain.Task

	result := db.First(&task, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Задача не найдена"})
		return
	}

	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	db.Save(&task)

	c.JSON(200, task)
}
