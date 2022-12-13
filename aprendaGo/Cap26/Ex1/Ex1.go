// - Crie um package "cachorro".
//     - Este package deverá exportar uma função Idade, que toma como parâmetro um número de anos e retorna a idade equivalente em anos caninos. (1 ano humano → 7 anos caninos)
//     - Documente seu código com comentários, e utilize a função Idade na sua função main.
// - Rode seu programa para verificar se ele funciona.
// - Rode um local server com godoc e leia sua documentação.
package main

import "fmt"

// Idade calcula quanto anos nós temos em anos caninos
func Idade(x int) int {
    return x * 7
}

func main() {

    fmt.Println(Idade(10))
    fmt.Println("Hello World")

}