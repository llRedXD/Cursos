// - v, ok := ←chan
// - Se receber valor: v, true
// - Canal fechado, nada, etc.: zero v, false
// - Agora vamos resolver o problema do exercício anterior usando comma ok.
// - Código: https://github.com/vkorbes/aprendago/blob/master/código/21_canais/06_exercício_anterior/main.go
package main

import "fmt"

func main() {

    canal := make(chan int)

    go func() {
        canal <- 42
        close(canal)
    }()
    
    v, ok := <- canal
    fmt.Println(v, ok) // Verdadeiro
    
    v, ok = <- canal
    fmt.Println(v, ok) // Falso
}