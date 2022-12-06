// - If, else.
// - If, else if, else.
// - If, else if, else if, ..., else.
package main

import "fmt"

func main() {
	x := 101
	if x < 100 {
		fmt.Println("x é menor que 100")
	} else if x == 10 {
		fmt.Println("x é igual a 10")
	} else {
		fmt.Println("x é maior que 100")
	}
}