package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) // Zero value é false
	x = true
	fmt.Println(x) // Valor atribuído
	y := 1 == 10
	z := 1 > 2
	fmt.Println(y)
	fmt.Println(z)
}