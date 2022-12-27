package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"github.com/llRedXD/go5/models"
)

var (
	Db *gorm.DB
	err error
)

func ConectDb() {
	dsn := "host=localhost user=root password=root dbname=root port=5433 sslmode=disable"
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println("Erro:", err)
	}
	Db.AutoMigrate(&models.Aluno{})
}
