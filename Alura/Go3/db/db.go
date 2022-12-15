package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
)

func ConectDb() *sql.DB {
	con := "user=postgres dbname=alura_loja sslmode=disable host=localhost password=123456"
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Println("Erro:", err)
	}
	return db
}
