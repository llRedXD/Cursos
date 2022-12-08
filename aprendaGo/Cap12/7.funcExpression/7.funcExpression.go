// - f := func(p params){ ... }
// - f()
package main

import "fmt"

func main() {
	x := 30

	y := func (x int) int {
		// fmt.Println(x * 10)
		return x * 10
	}

	fmt.Println(x, "Vezes 10 é", y(x))


}