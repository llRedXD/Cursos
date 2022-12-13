// Recover é uma função que permite que um programa continue a execução após um erro. Ele corta o fluxo de execução do programa e retorna o controle para a função que chamou a função que causou o erro. A função recover() retorna o valor que foi passado para a função panic(). Se a função panic() não foi chamada, a função recover() retorna nil. A função recover() só pode ser chamada dentro de uma função defer. Se a função recover() for chamada fora de uma função defer, ela não fará nada e retornará nil.
package main

import "fmt"

func main() {
    f()
    fmt.Println("Retornando normalmente de F.")
}

func f() {
    defer func() {
        if r := recover(); r != nil {
            fmt.Println("Recuperado de ", r)
        }
    }()
    fmt.Println("Chamando G.")
    g(0)
    fmt.Println("Retornando normalmente de G.")
}

func g(i int) {
    if i > 3 {
        fmt.Println("Panico!")
        panic(fmt.Sprintf("%v", i))
    }
    defer fmt.Println("Defer em G.", i)
    fmt.Println("Imprimindo em G.", i)
    g(i + 1)
}

