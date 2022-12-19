package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConectDb() *gorm.DB {
	dsn := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Erro:", err)
	}
	return db
}
