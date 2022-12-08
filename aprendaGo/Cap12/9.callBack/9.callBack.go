// - Primeiro veja se você entende isso: https://play.golang.org/p/QkAtwMZU-z
// - Callback é passar uma função como argumento.
// - Exemplo:
//     - Criando uma função que toma uma função e um []int, e usa somente os números pares como argumentos para a função.
//     - Go Playground:
// - Desafio: Crie uma função no programa acima que utilize somente os números ímpares.
package main

import "fmt"

func main() {
	fmt.Println(somenteImpar(soma,[]int{10, 11, 20,25,14,8541,6984,5,49,81,65,4165,4,56,4,654,65,42}...))
}

func soma(x ... int) int {
	t := 0
	for _, i := range x{
		t += i
	}
	return t
}

func somenteImpar(f func (y ... int) int,x ... int) int {
	slice := []int{}
	for _,v := range x {
		if v % 2 != 0 {
			slice = append(slice, v)
		}
	}
	n := soma(slice...)
	return n
	
}