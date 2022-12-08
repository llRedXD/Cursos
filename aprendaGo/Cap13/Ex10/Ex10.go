// - Demonstre o funcionamento de um closure.
// - Ou seja: crie uma função que retorna outra função, onde esta outra função faz uso de uma variável alem de seu scope
package main

import "fmt"

func main() {

    a := somaClosure(5)
    fmt.Println(a())
    fmt.Println(a())
    fmt.Println(a())

    b := somaClosure(2)
    fmt.Println(b())
    fmt.Println(b())
    fmt.Println(b())
}

func somaClosure(y int) func() int{
    x := y
    return func () int {
        x++
        return x
    }    
}
