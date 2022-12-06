// - Switch:
//   - pode avaliar uma expressão
//   - switch statement == case (value)
//   - default switch statement == true (bool)
//   - não há fall-through por padrão
//   - criando fall-through
//   - default
//   - cases compostos
package main

import "fmt"

func main() {
	fmt.Println("Switch")
	
	x := 500

	switch  {
		case x < 10:
			fmt.Println("x é menor que 10")
		case x > 100:
			fmt.Println("x é maior que 100")
		case x == 50:
			fmt.Println("x é igual a 50")
	}

	quemTaNoEscritorio := ""

	switch quemTaNoEscritorio {
		case "João":
			fmt.Println("João está no escritório")
			fallthrough // força o próximo case a ser executado
		case "Maria":
			fmt.Println("Maria está no escritório")
		case "Pedro":
			fmt.Println("Pedro está no escritório")
		case "José":
			fmt.Println("José está no escritório")
		default:
			fmt.Println("Ninguém está no escritório")
	}

	quemTaNoEscritorio = "Maria" 
	switch quemTaNoEscritorio{
		case "João", "Maria":
			fmt.Println("João ou Maria está no escritório")
		case "Pedro", "José":
			fmt.Println("Pedro ou José está no escritório")
	}
}