// - Alem da goroutine principal, crie duas outras goroutines.
// - Cada goroutine adicional devem fazer um print separado.
// - Utilize waitgroups para fazer com que suas goroutines finalizem antes de o programa terminar.
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func main() {
    wg.Add(2)
    go contador()
    go printaOi()
    wg.Wait()
}

func contador() {
    for i := 0; i < 10; i++ {
        fmt.Println(i)
        runtime.Gosched()
    }
    wg.Done()
}

func printaOi() {
    for i := 0; i < 10; i++ {
        fmt.Println("Oi")
        runtime.Gosched()
    }
    wg.Done()
}
