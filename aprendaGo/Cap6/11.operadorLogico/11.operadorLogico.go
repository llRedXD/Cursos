// - && (AND)
// - || (OR)
// - !
// - Go Playground: https://play.golang.org/p/MFwrt93xlc
// - Qual o resultado de fmt.Println...
//   - true && true
//   - true && false
//   - true || true
//   - true || false
//   - !true
package main

import "fmt"

func main() {
	x := 2
	y := 3

	if (x == 2) || (x ==3) {
		fmt.Println("x é igual a 2 ou 3")
	} 

	if (x == 2) && (y == 3) {
		fmt.Println("x é igual a 2 e y é igual a 3")
	}

}