    // - gofunc com for loop com send e close(chan)
    // - recebe com range chan
package main

import "fmt"

func main() {

    c := make(chan int)
    go meuLoop(10, c)
    for v := range c {
        fmt.Println(v)
    }

}

func meuLoop(t int,s chan<- int) {
    for i := 0; i < t; i++{
        s <- i
    }
    close(s)
}
