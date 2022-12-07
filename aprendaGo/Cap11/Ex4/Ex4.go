// - Crie e use um struct anônimo.
// - Desafio: dentro do struct tenha um valor de tipo map e outro do tipo slice.
package main

import "fmt"

func main() {


	
	x := struct {
		nome  string
		sobrenomes []string
		parente  map[string]string
	}{
		nome: "Gabriel",
		sobrenomes: []string{"Henrique", "Oliveira", "Neira"},
		parente: map[string]string {"Pai":"Jefferson"},
	}

	fmt.Println(x)
}