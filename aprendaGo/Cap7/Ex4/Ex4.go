// - Crie um loop utilizando a sintaxe: for {}
// - Utilize-o para demonstrar os anos desde que você nasceu.
package main

import "fmt"

func main() {
	x := 2004
	for {
		if x > 2022 {
			break
		}
		fmt.Println(x)
		x++
	}
}