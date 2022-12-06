// - If: bool
// - If: o operador não → "!"
// - If: declaração de inicialização
package main

import "fmt"

func main() {
	x := 10

	if x < 100 {
		fmt.Println("x é menor que 100")
	}

	if !(x > 100) { // ! sinal de operação negativa
		fmt.Println("x é menor que 100")
	}

	if y := 200; y > 199 { // declaração de inicialização
		fmt.Println("y é maior que 199")
	}
}