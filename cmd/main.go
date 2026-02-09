package main

import (
	"fmt"

	"github.com/dsolyakin/task-tracker/domain"
	"github.com/dsolyakin/task-tracker/internal/http"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	dsn := "host=localhost user=postgres password=pass123 dbname=tasktracker_db port=5432 sslmode=disable"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&domain.Task{})
	if err != nil {
		panic("Не удалось выполнить миграцию базы данных")
	}

	fmt.Println("Миграция выполнена успешно")

	r := gin.Default()
	http.InitRoutes(r, db)
	r.Run()
}
