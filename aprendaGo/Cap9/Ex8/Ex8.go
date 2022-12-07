// - Crie um map com key tipo string e value tipo []string.
//   - Key deve conter nomes no formato sobrenome_nome
//   - Value deve conter os hobbies favoritos da pessoa
//
// - Demonstre todos esses valores e seus índices.
package main

import "fmt"

func main() {
	x := map[string]string {
		"neira_gabriel": "jogos",
		"neira_isabelly": "make",
		"neira_lilian" : "dorama",
	}

	for key, value := range	x {
		fmt.Println(key, value)
	}
}