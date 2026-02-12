package domain

type User struct {
	ID        uint   `json:"id"`
	FirstName string `json:"first_name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"-" gorm:"not null"`
}
