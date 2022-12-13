// - Canais são o Jeito Certo® de fazer sincronização e código concorrente.
// - Eles nos permitem trasmitir valores entre goroutines.
// - Servem pra coordenar, sincronizar, orquestrar, e buffering.
// - Na prática:
//     - make(chan type, b)
// - Canais bloqueiam:
//     - Eles são como corredores em uma corrida de revezamento
//     - Eles tem que "passar o bastão" de maneira sincronizada
//     - Se um corredor tentar passar o bastão pro próximo, mas o próximo corredor não estiver lá...
//     - Ou se um corredor ficar esperando receber o bastão, mas ninguem entregar...
//     - ...não dá certo.
// - Exemplos:
//     - Poe um valor num canal e faz um print. Block.
//         - Código acima com goroutine.
//         - Ou com buffer. Via de regra: má idéia; é legal em certas situações, mas em geral é melhor sempre passar o bastão de maneira sincronizada.
// - Interessante: ref/spec → types
package main

import "fmt"

func main() {
    canal := make(chan int)
    go func() {
        canal <- 42
    }()
    fmt.Println("Com Go routine",<-canal)
    // Buffer
    canal1 := make(chan int, 1)
    canal1 <- 42
    // canal <- 43 Block
    fmt.Println("Com buffer:", <-canal1)

}


