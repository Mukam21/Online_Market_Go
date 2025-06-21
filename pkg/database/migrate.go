package database

import (
	"Online_market/pkg/models"
	"log"
)

func AutoMigrate() {
	db := GetDB()
	err := db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
	)
	if err != nil {
		log.Fatalf("Ошибка миграции: %v", err)
	}
	log.Println("Миграция прошла успешно")
}
