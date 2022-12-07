// - Utilizando o exercício anterior, adicione uma entrada ao map e demonstre o map inteiro utilizando range.
package main

import "fmt"

func main() {
	x := map[string]string {
		"neira_gabriel": "jogos",
		"neira_isabelly": "make",
		"neira_lilian" : "dorama",
	}

	x["neira_jefferson"] = "futebol"

	for key, value := range	x {
		fmt.Println(key, value)
	}
}