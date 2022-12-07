// - Crie um tipo "pessoa" com tipo subjacente struct, que possa conter os seguintes campos:
//   - Nome
//   - Sobrenome
//   - Sabores favoritos de sorvete
//
// - Crie dois valores do tipo "pessoa" e demonstre estes valores, utilizando range na slice que contem os sabores de sorvete.
package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	saboresFavoritoSorvete []string
}

func main() {
	pessoa1 := pessoa{
		nome: "Gabriel",
		sobrenome: "Neira",
		saboresFavoritoSorvete: []string{"Chocolate", "Flocos"},
	}
	pessoa2 := pessoa {
		nome: "Lilian",
		sobrenome: "Neira",
		saboresFavoritoSorvete: []string{"Ovomaltine", "Creme"},
	}
	
	fmt.Printf("Nome: %v, Sobrenome: %v\n", pessoa1.nome, pessoa1.sobrenome)

	for _, v := range pessoa1.saboresFavoritoSorvete {
		fmt.Println(v)
	}
	fmt.Printf("Nome: %v, Sobrenome: %v\n", pessoa2.nome, pessoa2.sobrenome)

	for _, v := range pessoa2.saboresFavoritoSorvete {
		fmt.Println(v)
	}
	
}