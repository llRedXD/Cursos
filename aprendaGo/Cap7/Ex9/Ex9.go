// - Crie um programa que utilize a declaração switch, onde o switch statement seja uma variável do tipo string com identificador "esporteFavorito".
package main

import "fmt"

func main() {
	esporteFavorito := "Futebol"

	switch esporteFavorito {
		case "Basquete":
			fmt.Println("Amo Basquete")
		case "Vólei":
			fmt.Print("Amo Vólei")
		case "Futebol":
			fmt.Println("Amo Futebol")
	}
}
