// - Structs dentro de structs dentro de structs.
// - Exemplo: um corredor de fórmula 1 é uma pessoa (nome, sobrenome, idade) e tambem um competidor (nome, equipe, pontos).
package main

import "fmt"

type pessoa struct {
	nome string
	idade int
}

type profissional struct {
	pessoa 
	titulo string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome: "Isa",
		idade: 14,
	}

	pessoa2 := profissional{
		pessoa: pessoa{
			nome: "Gabriel",
			idade: 18,
		},
		titulo: "Programador",
		salario: 1800,	
	}

	// Podemos declarar uma struct de uma forma mais simplificada, sem precisar declarar os campos.
	// pessoa3 := pessoa{"João", 20}
	// pessoa4 := profissional{"João", 20, "Programador", 1800}

	fmt.Println(pessoa1)
	fmt.Println(pessoa2)
	// Para acessar os campos de uma struct dentro de outra, basta usar o nome da struct e o ponto.
	fmt.Println(pessoa2.nome)
	
	
}