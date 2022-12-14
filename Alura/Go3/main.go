package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"log"
	"net/http"

	_ "github.com/lib/pq"
)

type produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func conectDb() *sql.DB {
	con := "user=postgres dbname=alura_loja sslmode=disable host=localhost password=123456"
	db, err := sql.Open("postgres", con)
	if err != nil {
		log.Println("Erro:", err)
	}
	return db
}

func main() {

	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	db := conectDb()
	defer db.Close()

	selectProduto, err := db.Query("select * from produtos")
	if err != err {
		log.Println("Erro:", err)
	}
	p := produto{}
	produtos := []produto{}

	for selectProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectProduto.Scan(&id,&quantidade,&nome,&descricao,&preco)
		if err != nil {
			log.Println("Erro:", err)
		}
		p.Id = id
		p.Quantidade = quantidade
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco

		produtos = append(produtos, p)
	}
	fmt.Println(produtos)

	temp.ExecuteTemplate(w, "Index", produtos)
}
