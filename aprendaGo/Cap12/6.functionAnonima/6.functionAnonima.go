// - Anonymous self-executing functions → Funções anônimas auto-executáveis.
// - func(p params) { ... }()
// - Vamos ver bastante quando falarmos de goroutines.
package main

import "fmt"

func main() {
	x := 10

	func (x int) {
		fmt.Println(x * 10)
	}(x)


}