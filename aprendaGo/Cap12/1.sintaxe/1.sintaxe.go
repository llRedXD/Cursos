// - Qual a utilidade de funções?
//   - Abstrair funcionalidade
//   - Reutilização de código
//
// - func (receiver) identifier(parameters) (returns) { code }
// - A diferença entre parâmetros e argumentos:
//   - Funções são definidas com parâmetros
//   - Funções são chamadas com argumentos
//
// - Tudo em Go é pass by value.
//   - Pass by reference, pass by copy, ... não.
//
// - Parâmetro pode ser ...variádico.
// - Exemplos:
//   - Função básica.
//   - Go Playground: https://play.golang.org/p/FebJblBenP
//   - Função que aceita um argumento.
//   - Go Playground:
//     https://play.golang.org/p/CE6Ij3U4QB
//   - Função com retorno.
//   - Go Playground: https://play.golang.org/p/gKxwYe6btP
//   - Função com múltiplos retornos e parâmetro variádico.
//   - Go Playground: https://play.golang.org/p/OcQ1wXwM2c
//   - Mais um: https://play.golang.org/p/8wc2TA9xH_
package main

import "fmt"

func main() {
	s := "abcasad"

	fmt.Println(len(s)) // Len é uma função que retorna o tamanho de algo. Quando chamamos uma função passamos argumentos.
	basica() 
	argumento(1) 
	valor := soma(1, 2)
	fmt.Println(valor)
	fmt.Println(variadico(1,2,3,4,5,6,7,8,9,10))
}

func basica () {
	fmt.Println("Função básica")
}

func argumento (a int) { // Quando definimos uma função, definimos parâmetros.
	fmt.Println(a)
}

func soma(a,b int) int { // Quando definimos uma função, definimos parâmetros.
	return a + b
}

func variadico(a ...int) int { // Parâmetro pode ser variádico.
	total := 0
	for _, v := range a {
		total += v
	}
	return total
}