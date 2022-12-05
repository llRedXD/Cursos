// - int vs. float: Números inteiros vs. números com frações.
// - golang.org/ref/spec → numeric types
// - Integers:
//     - Números inteiros
//     - int & uint → “implementation-specific sizes”
//     - Todos os tipos numéricos são distintos, exceto:
//         - byte = uint8
//         - rune = int32 (UTF8)
//         (O código fonte da linguagem Go é sempre em UTF-8).
//     - Tipos são únicos
//         - Go é uma linguagem estática
//         - int e int32 não são a mesma coisa
//         - Para "misturá-los" é necessário conversão
//     - Regra geral: use somente int
// - Floating point:
//     - Números racionais ou reais
//     - Regra geral: use somente float64
// - Na prática:
//     - Defaults com :=
//     - Tipagem com var
//     - Dá pra colocar número com vírgula em tipo int?
//     - Overflow
//     - Go Playground: https://play.golang.org/p/dt2x1ies5b
// - “implementation-specific sizes”? Runtime package. Word.
//     - GOOS
//     - GORUNTIME

package main

import (
	"fmt"
	"runtime"
)

func main() {
	a := "e"
	b := "é"
	c := "\u9999"
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)

	practice()
	saberComputador()
}

var g = 40
var h = 40.0

func practice() {
	x := 10
	y := 10.0
	fmt.Printf("%v, %T\n", x, x)
	fmt.Printf("%v, %T\n", y, y)
	fmt.Printf("%v, %T\n", g, g)
	fmt.Printf("%v, %T\n", h, h)
}

func saberComputador() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}