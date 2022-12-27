package db

import (
	"log"

	"github.com/llRedXD/Go6/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	Db  *gorm.DB
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
