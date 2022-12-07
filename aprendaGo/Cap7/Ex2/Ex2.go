//   - Põe na tela: O unicode code point de todas as letras maiúsculas do alfabeto, três vezes cada.
//   - Por exemplo:
//     65
//     U+0041 'A'
//     U+0041 'A'
//     U+0041 'A'
//     66
//     U+0042 'B'
//     U+0042 'B'
//     U+0042 'B'
//     ...e por aí vai.
package main

import "fmt"

func main() {
	for y := 65; y < 91; y++ {
		fmt.Println(y)
		for z := 0;z < 3;z++ {
			fmt.Printf("\t%#U\n", y)
		}
	}
}