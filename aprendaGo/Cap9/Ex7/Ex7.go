// - Crie uma slice contendo slices de strings ([][]string). Atribua valores a este slice multi-dimensional da seguinte maneira:
//   - "Nome", "Sobrenome", "Hobby favorito"
//
// - Inclua dados para 3 pessoas, e utilize range para demonstrar estes dados.
package main

import "fmt"

func main() {
	dados := [][]string {
		{"Gabriel", "Neira", "Jogos"},
		{"Isabelly", "Neira", "Anime"},
		{"Lilian", "Neira","Dorama"},
		{"Jefferson", "Neira", "Futebol"},
	}

	for i := range dados {
		for _, t := range dados[i] {
			fmt.Print(t, " " )
		}
		fmt.Print("\n")
	}
	
}