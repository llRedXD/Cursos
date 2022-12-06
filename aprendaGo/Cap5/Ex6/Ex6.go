// - Utilizando iota, crie 4 constantes cujos valores sejam os próximos 4 anos.
// - Demonstre estes valores.

package main

import "fmt"

const (
	a = iota + 2023
	b 
	c 
	d 
)

func main() {
	fmt.Println(a, b, c, d)
}
