package main

import "fmt"

// - Declaração vs. inicialização vs. atribuição de valor. Variáveis: caixas postais.
// - Declaração: criar uma caixa postal.
// - Inicialização: colocar um endereço na caixa postal (Primeira atribuição de valor).
// - Atribuição de valor: colocar uma carta na caixa postal.
// - O que é valor zero?
// - Os zeros:
//     - Ints: 0
//     - Floats: 0.0
//     - Booleans: false
//     - Strings: ""
//     - Pointers, Functions, Interfaces, Slices, Channels, Maps: nil
// - Use := sempre que possível.
// - Use var para package-level scope.

var x int // Declaração

func main() {
	// Mas podemos usar o := para fazer a declaração e inicialização ao mesmo tempo, por isso operador curto de declaração.
	y := 42 // Declaração e inicialização
	fmt.Println(y)
	x = 42 // Inicialização 
	fmt.Println(x)
	x = 99 // Atribuição de valor
	fmt.Println(x)
	tipos()
}

func tipos() {
	// Declaração e inicialização
	var a int
	var b string
	var c float64
	var d bool
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
	fmt.Printf("%v, %T\n", d, d)
}