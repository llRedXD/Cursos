// - Crie uma função que retorna uma função.
// - Atribua a função retornada a uma variável.
// - Chame a função retornada.
package main

import "fmt"

func vezesDez() func(x int) int {
    return func(x int) int {
        return x * 10
    }
}

func main() {

    fmt.Println(vezesDez()(5))
}