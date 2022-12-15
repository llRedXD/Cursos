package models

import (
	"log"

	"github.com/llRedXD/Go3/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func PegaTodosOsProdutos() []Produto {
	db := db.ConectDb()
	defer db.Close()

	selectProdutos, err := db.Query("select * from produtos")
	if err != err {
		log.Println("Erro:", err)
	}
	p := Produto{}
	produtos := []Produto{}

	for selectProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err := selectProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
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
	return produtos
}

func CriaProduto(nome, descricao string, quantidade int, preco float64) {
	db := db.ConectDb()
	defer db.Close()

	selectProduto, err := db.Prepare("insert into produtos (nome, descricao, quantidade, preco) values ($1, $2, $3, $4)")
	if err != nil {
		log.Println("Erro:", err)
	}

	_, erro := selectProduto.Exec(nome, descricao, quantidade, preco)
	if erro != nil {
		log.Println("Erro:", erro)
	}
}

func Delete(id string) {
	db := db.ConectDb()
	defer db.Close()

	deleteProduto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		log.Println("Erro:", err)

	}

	_, erro := deleteProduto.Exec(id)
	if erro != nil {
		log.Println("Erro:", erro)
	}
}

func PegaProduto(ids string) Produto {
	db := db.ConectDb()
	defer db.Close()

	selectProduto, err := db.Query("select * from produtos where id =" + ids)
	if err != err {
		log.Println("Erro:", err)
	}

	p := Produto{}

	for selectProduto.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		erro := selectProduto.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if erro != nil {
			log.Println("Erro:", erro)
		}
		p.Id = id
		p.Quantidade = quantidade
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
	}

	return p
}
