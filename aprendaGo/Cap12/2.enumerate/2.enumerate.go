// - Quando temos uma slice, podemos passar os elementos individuais através "deste..." operador.
// - Exemplos:
//   - Desenrolando uma slice de ints com como argumento para a função "soma" anterior
package main

import "fmt"

func main() {
	fmt.Println(soma(1,2,3,4,5,6,7,8,9,10))
	s := []int{1,2,3,4,5,6,7,8,9,10}
	fmt.Println(soma(s...)) // Desenrolando uma slice de ints com como argumento para a função "soma" anterior
}

func soma(a ...int) int { // Parâmetro variádico (...) tem que ser o ultimo parâmetro sempre
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}