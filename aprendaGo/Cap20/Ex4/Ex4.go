// - Utilize mutex para consertar a condição de corrida do exercício anterior.
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
    totalDeGoRoutines := 10
    
    var wg sync.WaitGroup
    var mu sync.Mutex
    wg.Add(totalDeGoRoutines)
    for i := 0; i < totalDeGoRoutines; i++ {
        go func() {
            mu.Lock()
            v := cont
            runtime.Gosched()
            v++
            cont = v
            mu.Unlock()
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