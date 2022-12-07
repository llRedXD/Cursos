// - Utilizando a solução anterior, adicione as opções else if e else.
package main

import "fmt"

func main() {
	x := 3
	if x == 1 {
		fmt.Println("Se x for igual 1 vai ativar esse if")
	}  else  if x == 2{
		fmt.Println("Se x for igual 2 vai ativar esse else if")
	} else {
		fmt.Println("se x for diferente de 1 e 2 vai ativar o else")
	}
}