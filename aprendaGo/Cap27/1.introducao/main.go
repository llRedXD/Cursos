// - gofmt: formata o código
// - go vet: correctness → procura constructs suspeitos
// - golint: suggestions → procura coding style ruim
package main

import "fmt"

func multiplica(a, b int) int {
	return a * b
}

func soma(a, b int) int {
	return a + b
}

func main() {

	s := soma(1, 2)
	x := multiplica(2, 2)
	fmt.Println(s,x)
}

