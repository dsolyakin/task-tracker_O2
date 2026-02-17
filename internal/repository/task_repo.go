package repository

import (
	"github.com/dsolyakin/task-tracker/domain"
	"gorm.io/gorm"
)

type TaskRepository interface {
	Create(task domain.Task) error
	GetByID(id uint) (*domain.Task, error)
}

type GormTaskRepository struct {
	db *gorm.DB
}

// КОНСТРУКТОР
// 1. Название: принято называть New + ИмяСтруктуры
// 2. Вход: Мы просим дать нам "инструмент" (database *gorm.DB)
// 3. Выход: Мы обещаем отдать готовую, "заряженную" структуру
func NewGormTaskRepository(database *gorm.DB) *GormTaskRepository {

	// 1. Создаем структуру
	// 2. Кладем в её "ячейку" db нашу базу
	// 3. Ставим значок & (берем адрес)
	// 4. Выкидываем этот адрес наружу через return
	return &GormTaskRepository{
		db: database,
	}
}
