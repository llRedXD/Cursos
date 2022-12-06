// - Operação módulo: %
// - For: break
// - For: continue
package main

import "fmt"

func main() {
	for x := 0; x < 21; x++ {
		if x%2 != 0 {
			continue
		}
		fmt.Println(x)
	}
}