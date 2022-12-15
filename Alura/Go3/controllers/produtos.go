package controllers

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/llRedXD/Go3/models"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	produtos := models.PegaTodosOsProdutos()

	temp.ExecuteTemplate(w, "Index", produtos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		quantidade := r.FormValue("quantidade")
		preco := r.FormValue("preco")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("Erro ao converter preço")
		}

		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("Erro ao converter quantidade")
		}

		models.CriaProduto(nome, descricao, quantidadeConvertida, precoConvertido)
	}
	http.Redirect(w, r, "/", int(301))
}

func Delete(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	models.Delete(idProduto)
	http.Redirect(w, r, "/", int(301))
}

func Edit(w http.ResponseWriter, r *http.Request) {
	produto := models.PegaProduto(r.URL.Query().Get("id"))
	temp.ExecuteTemplate(w, "Edit", produto)
}