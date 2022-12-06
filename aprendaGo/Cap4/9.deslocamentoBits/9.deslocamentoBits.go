// Deslocamento de bits é quando deslocamos dígitos binários para a esquerda ou direita.
package main

import "fmt"


const (
	_ = iota
	kb = 1 << (iota * 10) // 1 << (1 * 10)
	mb // 1 << (2 * 10)
	gb // 1 << (3 * 10)
	tb // 1 << (4 * 10)
)


func main() {
	x := 100
	t := x << 1 // Deslocamento de bits para a esquerda. Multiplica por 2 elevado ao número de deslocamentos.
	y := x >> 1 // Deslocamento de bits para a direita. Divide por 2 elevado ao número de deslocamentos.
	fmt.Printf("%b <- %b -> %b\n", t, x, y)
	fmt.Printf("%v <- %v -> %v\n", t, x, y)
	fmt.Printf("%v KB\n", kb)
	fmt.Printf("%v MB\n", mb)
	fmt.Printf("%v GB\n", gb)
	fmt.Printf("%v TB\n", tb)
}