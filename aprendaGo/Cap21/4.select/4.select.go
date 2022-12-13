// - Select é como switch, só que pra canais, e não é sequencial.
// - "A select blocks until one of its cases can run, then it executes that case. It chooses one at random if multiple are ready." — https://tour.golang.org/concurrency/5
// - Na prática:
// - Exemplo 1:
//   - Duas go funcs enviando X/2 numeros cada uma pra um canal
//   - For loop X valores, select case ←x

// - Exemplo 2:
//   - Func 1 recebe X valores de canal, depois manda qualquer coisa pra chan quit
//   - Func 2 for infinito, select: case envia pra canal, case recebe de quit
//
// - Exemplo 3:
//   - Chans par, ímpar, quit
//   - Func send manda números pares pra um canal, ímpares pra outro, e fecha/quit
//   - Func receive é um select entre os três canais, encerra no quit
//   - Problema!
package main

import (
	"fmt"
)


func main() {
    par := make(chan int)
    impar := make(chan int)
    quit := make(chan bool)

    go send(par, impar,quit)

    receive(par, impar, quit)
}

func send(par chan int, impar chan int, quit chan bool) {
    for i := 0; i < 50; i++ {
        if i % 2 == 0 {
            par <- i
        } else {
            impar <- i
        }
    }
    close(par)
    close(impar)
    quit <- true
}

func receive(par chan int, impar chan int, quit chan bool)  {
    for {
        select {
        case v, ok := <- par:
            if ok {
                fmt.Println("O numero",v, " é par")
            }
        case v, ok := <- impar:
            if ok{
                fmt.Println("O numero",v, " é impar")
            }
        case <- quit:
            return
            
        }
    }
}

