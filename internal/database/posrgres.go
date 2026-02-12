package database

import (
	"fmt"

	"github.com/dsolyakin/task-tracker/domain"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(dsn string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&domain.Task{}, &domain.Category{}, &domain.Tag{})
	if err != nil {
		panic("Не удалось выполнить миграцию базы данных")
	}
	fmt.Println("Миграция выполнена успешно")

	return db
}
