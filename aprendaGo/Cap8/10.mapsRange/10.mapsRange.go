// - Range: for k, v := range map { }
// - Reiterando: maps não tem ordem e um range usará uma ordem aleatória.
// - delete(map, key)
// - Deletar uma key não-existente não retorna erros!
package main

import "fmt"

func main() {
	x := map[string]int{
		"joao":  123456,
		"maria": 654321,
		"pedro": 987654,
	}
	fmt.Println(x)

	total := 0

	for _, v := range x {
		fmt.Println(v)
		total += v
	}

	fmt.Println(total)

	delete(x, "joao")

	fmt.Println(x)

}