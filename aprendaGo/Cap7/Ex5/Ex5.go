// - Demonstre o resto da divisão por 4 de todos os números entre 10 e 100
package main

import "fmt"

func main() {
	for x := 10; x <= 100; x++{
		y := x % 4
		fmt.Printf("%v / %v = %v\n", x, 4, y)
	}
}