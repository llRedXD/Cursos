// - Pode avaliar tipos
// - Pode haver uma expressão de inicialização
package main

import "fmt"

func main() {
	var x interface{}

	x = "10.23"

	switch x.(type) {
		case int:
			fmt.Println("x é um inteiro")
		case float64:
			fmt.Println("x é um float64")
		case string:
			fmt.Println("x é uma string")
		case bool:
			fmt.Println("x é um booleano")
		
	}

	switch y := 1; y {
		case 1:
			fmt.Println("y é igual a 1")
		case 2:
			fmt.Println("y é igual a 2")
		case 3:
			fmt.Println("y é igual a 3")
	}

}