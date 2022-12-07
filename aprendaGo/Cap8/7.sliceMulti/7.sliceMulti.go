// - Slices multi-dimensionais são slices que contem slices.
// - São como planilhas.
// - [][]type
package main

import "fmt"

func main() {
	ss := [][]int { // Slice multi-dimensional pode ter quantas slices internas quiser.
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(ss) 
	fmt.Println(ss[0]) // Quando trabalhamos com slices multi-dimensionais, precisamos especificar o índice de cada slice.
	fmt.Println(ss[1][0]) // Assim podemos acessar o valor de cada slice dentro da slice multi-dimensional.
}