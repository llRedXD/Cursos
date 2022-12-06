// - For
//   - Repetição hierárquica
//   - Exemplos: relógio, calendário
package main

import "fmt"

func main() {
	for hora := 0; hora < 13; hora++ {
		fmt.Println("Hora:", hora)
		for minuto :=0; minuto < 60; minuto++ {
			fmt.Printf("%v ", minuto)
		}
		fmt.Println("")
	}
}