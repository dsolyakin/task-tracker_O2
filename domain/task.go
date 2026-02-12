package domain

type Task struct {
	ID          int       `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CategoryID  *uint     `json:"category_id" gorm:"default:null"`
	Category    *Category `json:"category" gorm:"foreignKey:CategoryID"`
	Tags        []Tag     `json:"tags" gorm:"many2many:task_tags;"`
}
