package main

import "fmt"

func main() {
	// O código vai começar a ser executado a partir daqui

	// Temos os tipos primitivos de variáveis
	x := 10 // int
	y := "string" // string
	z := true // bool
	
	numeroDeBytes, erros := fmt.Println(x, y, z)
	
	// A função Println vai imprimir na tela as strings que forem passadas como parâmetro, além de adicionar um espaço entre elas e um caractere de nova linha no final, pode mostrar também numero de bytes escritos e um erro
    
	fmt.Println(numeroDeBytes, erros) // Se as variáveis não forem usadas, o compilador vai dar um erro
	
	// Se eu não quiser usar uma variável, posso usar o _ para ignorá-la

	_, erros = fmt.Println(x, y, z)

	fmt.Println(erros)

	// O código vai terminar de ser executado aqui
}