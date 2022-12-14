package main

import (
	"fmt"
	"log"
)

func main() {
	nome := "Gabriel"
	versao := 2.1
	fmt.Printf("Bem vindo %v\nO sistema esta na versão %v\n",nome, versao)

	fmt.Println("1- Inciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("3- Sair")

	comando := 0
	_,err := fmt.Scan(&comando)
	if err != nil{
		log.Fatal("Valor Digitado é invalido")
	}
	if comando == 1 {
		fmt.Println("Monitorando...")
	} else if comando == 2 {
		fmt.Println("Logs...")
	} else if comando == 0 {
		fmt.Println("Saindo")
	} else {
		fmt.Println("Opção não valida")
	}

	fmt.Println("Valor Recebido", comando)



}