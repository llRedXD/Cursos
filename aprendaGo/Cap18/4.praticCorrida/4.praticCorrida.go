// - Aqui vamos replicar a race condition mencionada no artigo anterior.
//     - time.Sleep(time.Second) vs. runtime.Gosched()
// - go help → go help build → go run -race main.go
package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main() { 
    fmt.Println("Numero de CPUs:",runtime.NumCPU())
    fmt.Println("Numero de Goroutines:",runtime.NumGoroutine())
    
    cont := 0
    totalDeGoRoutines := 1000
    
    var wg sync.WaitGroup
    wg.Add(totalDeGoRoutines)
    for i := 0; i < totalDeGoRoutines; i++ {
        go func() {
            v := cont
            runtime.Gosched()
            v++
            cont = v
            wg.Done()
            }()
    }
    
    fmt.Println("Numero de CPUs:",runtime.NumCPU())
    fmt.Println("Numero de Goroutines:",runtime.NumGoroutine())

    wg.Wait()
    fmt.Println(cont)
    
    fmt.Println("Numero de CPUs:",runtime.NumCPU())
    fmt.Println("Numero de Goroutines:",runtime.NumGoroutine())

}