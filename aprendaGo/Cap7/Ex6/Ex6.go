// - Crie um programa que demonstre o funcionamento da declaração if.
package main

import "fmt"

func main() {
	x := 1
	y := 2
	if x == 1 {
		fmt.Println("Se x for igual vai ativar o primeiro if")
	} 
	if y != 1 {
		fmt.Println("Mas como y é diferente de 1 vai ativar esse if")
	}
}