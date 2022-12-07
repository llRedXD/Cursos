// - Utilizando o exercício anterior, remova uma entrada do map e demonstre o map inteiro utilizando range.
package main

import "fmt"

func main() {
	x := map[string]string {
		"neira_gabriel": "jogos",
		"neira_isabelly": "make",
		"neira_lilian" : "dorama",
	}

	delete(x, "neira_lilian")

	for key, value := range	x {
		fmt.Println(key, value)
	}
}