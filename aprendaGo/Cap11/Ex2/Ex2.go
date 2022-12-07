// - Utilizando a solução anterior, coloque os valores do tipo "pessoa" num map, utilizando os sobrenomes como key.
// - Demonstre os valores do map utilizando range.
// - Os diferentes sabores devem ser demonstrados utilizando outro range, dentro do range anterior.
package main

import "fmt"

type pessoa struct {
	nome string
	sobrenome string
	saboresFavoritoSorvete []string
}

func main() {

	nomes := map[string]pessoa {}

	nomes["Neira"] = pessoa{
		nome: "Gabriel",
		sobrenome: "Neira",
		saboresFavoritoSorvete: []string{"Chocolate", "Flocos"},
	}

	nomes["Oliveira"] = pessoa{
		nome: "Lilian",
		sobrenome: "Oliveira",
		saboresFavoritoSorvete: []string{"Ovomaltine", "Creme"},
	}

	for _, v := range nomes {
		fmt.Println(v.nome, v.sobrenome)
		for _ ,s := range v.saboresFavoritoSorvete {
			fmt.Println("\t",s)
		}
	}
	
	
}