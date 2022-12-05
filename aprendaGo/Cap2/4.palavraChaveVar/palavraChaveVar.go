package main

import "fmt"

var y = 20 // Declaração de variável global

func main() {
	qualquerCoisa(10)
}

func qualquerCoisa(x int) {
	// Variáveis declaradas dentro de uma função só podem ser usadas dentro dela, mas variáveis declaradas fora de uma função podem ser usadas em qualquer lugar do código para isso usamos o var
	fmt.Println(x)
	fmt.Println(y)
}