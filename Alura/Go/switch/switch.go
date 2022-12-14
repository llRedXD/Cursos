package main

import (
	"fmt"
	"os"
)

func main() {
	Introducao()	
	Menu()
	comando := Opcoes()


	switch comando {
	case 1:
		fmt.Println("Monitorando")
	case 2:
		fmt.Println("Logs")
	case 0:
		fmt.Println("Saindo")
		os.Exit(0)
	default:
		fmt.Println("Opção Invalida")
		os.Exit(-1)
	}

	fmt.Println("Valor Recebido", comando)

}

func Menu() {
	fmt.Println("1- Inciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair")
}

func Introducao() {
	nome := "Gabriel"
	versao := 2.1
	fmt.Printf("Bem vindo %v\nO sistema esta na versão %v\n",nome, versao)
}

func Opcoes() int {
	comando := 0
	fmt.Scan(&comando)
	return comando
}