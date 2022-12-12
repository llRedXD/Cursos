// - Agora vamos resolver a race condition do programa anterior utilizando mutex.
// - Mutex é mutual exclusion, exclusão mútua.
// - Utilizando mutex somente uma thread poderá utilizar a variável contador de cada vez, e as outras deve aguardar sua vez "na fila."
// - Na prática:
//     - type Mutex
//         - func (m *Mutex) Lock()
//         - func (m *Mutex) Unlock()
// - RWMutex
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
    
    var mu sync.Mutex
    var wg sync.WaitGroup
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