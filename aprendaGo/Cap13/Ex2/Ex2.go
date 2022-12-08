// - Crie uma função que receba um parâmetro variádico do tipo int e retorne a soma de todos os ints recebidos;
// - Passe um valor do tipo slice of int como argumento para a função;
// - Crie outra função, esta deve receber um valor do tipo slice of int e retornar a soma de todos os elementos da slice;
// - Passe um valor do tipo slice of int como argumento para a função.
package main

import "fmt"

func main() {

    fmt.Println("Hello World")
    x := []int {10,45,7813,27498,16,9748,13,25,448,5263,41,654,32,1}
    y := []int {10,45,7813,27498,16,9748,13,25,44,1}
    fmt.Println(soma(x...))
    fmt.Println(somaSlice(y))
}

func soma(x ... int) int {
    total := 0
    for _, v := range x {
        total += v 
    }
    return total
}

func somaSlice(x []int) int {
    total := 0
    for _, v := range x {
        total += v
    }
    return total
}