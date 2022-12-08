// - Pode-se usar uma função como retorno de uma função
// - Declaração: func f() return
// - Exemplo: func f() func() int { [...]; return func() int{ return [int] } }
//   - ????: fmt.Println(f()())
package main

import "fmt"

func main() {
	x := retornaFunc()
	y := x(5)
	fmt.Println("5 Vezes 10 é:", y)

	fmt.Println(retornaFunc()(8)) // Forma simplificada de chamar uma função que retorna o valor de uma função dentro de outra função
}

func retornaFunc() func(x int) int {
	return func (i int) int {
		return i * 10
	}
}