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

type CategoryHandler struct {
	DB *gorm.DB
}

type TagHandler struct {
	DB *gorm.DB
}

func (t *TaskHandler) CreateTaskHandler(c *gin.Context) {
	var task domain.Task

	err := c.ShouldBindJSON(&task)
	if err != nil {
		fmt.Println("CreateTask. Ошибка парсинга json:", err)
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	result := t.DB.Create(&task)
	err = result.Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось сохранить в базу"})
		return
	}
	t.DB.Preload("Category").Preload("Tags").First(&task, task.ID)

	c.JSON(201, task)
}

func (t *TaskHandler) GetTaskListHandler(c *gin.Context) {
	var tasks []domain.Task

	result := t.DB.Preload("Category").Preload("Tags").Find(&tasks)
	err := result.Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось получить список задач"})
		return
	}

	c.JSON(200, tasks)
}

func (t *TaskHandler) DeleteTaskHandler(c *gin.Context) {
	id := c.Param("id")

	result := t.DB.Delete(&domain.Task{}, id)
	if result.RowsAffected == 0 {
		c.JSON(404, gin.H{"error": "Задача не найдена"})
		return
	}
	c.Status(204)
}

func (t *TaskHandler) GetTaskIdHandler(c *gin.Context) {
	id := c.Param("id")

	var task domain.Task

	result := t.DB.Preload("Category").Preload("Tags").First(&task, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Задача не найдена"})
		return
	}
	c.JSON(200, task)

}

func (t *TaskHandler) UpdateTaskHandler(c *gin.Context) {
	id := c.Param("id")

	var task domain.Task

	result := t.DB.First(&task, id)
	if result.Error != nil {
		c.JSON(404, gin.H{"error": "Задача не найдена"})
		return
	}

	err := c.ShouldBindJSON(&task)
	if err != nil {
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	t.DB.Save(&task)

	t.DB.Preload("Category").Preload("Tags").First(&task, task.ID)

	c.JSON(200, task)
}

func (cat *CategoryHandler) CreateCategoryHandler(c *gin.Context) {
	var category domain.Category

	err := c.ShouldBindJSON(&category)
	if err != nil {
		fmt.Println("CreateCategory. Ошибка парсинга json:", err)
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	result := cat.DB.Create(&category)
	err = result.Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось сохранить в базу"})
		return
	}
	c.JSON(201, category)
}

func (cat *CategoryHandler) GetCategoryListHandler(c *gin.Context) {
	var categories []domain.Category

	result := cat.DB.Find(&categories)
	err := result.Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось получить список категорий"})
		return
	}

	c.JSON(200, categories)
}

func (tg *TagHandler) CreateTagHandler(c *gin.Context) {
	var tag domain.Tag

	err := c.ShouldBindJSON(&tag)
	if err != nil {
		fmt.Println("CreateTag. Ошибка парсинга json:", err)
		c.JSON(400, gin.H{"error": "Неверный формат данных"})
		return
	}

	result := tg.DB.Create(&tag)
	err = result.Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось сохранить в базу"})
		return
	}
	c.JSON(201, tag)
}

func (tg *TagHandler) GetTagListHandler(c *gin.Context) {
	var tag []domain.Tag

	result := tg.DB.Find(&tag)
	err := result.Error
	if err != nil {
		c.JSON(500, gin.H{"error": "Не удалось получить список тегов"})
		return
	}

	c.JSON(200, tag)
}
