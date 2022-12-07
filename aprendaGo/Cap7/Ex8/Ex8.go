// - Crie um programa que utilize a declaração switch, sem switch statement especificado.
package main

import "fmt"

func main() {
	x := 20

	switch {
	case x == 10:
		fmt.Println("x igual a 10")
	case x == 15:
		fmt.Println("x igual a 15")
	case x == 20:
		fmt.Println("x igual a 20")
	}
}