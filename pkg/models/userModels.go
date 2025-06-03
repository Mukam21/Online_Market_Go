package models

type User struct {
	ID    uint `gorm:"primarykey"`
	Name  string
	Email string `gorm:"unique"`
}
