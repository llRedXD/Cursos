// - Utilizando goroutines, crie um programa incrementador:
//     - Tenha uma variável com o valor da contagem
//     - Crie um monte de goroutines, onde cada uma deve:
//         - Ler o valor do contador
//         - Salvar este valor
//         - Fazer yield da thread com runtime.Gosched()
//         - Incrementar o valor salvo
//         - Copiar o novo valor para a variável inicial
//     - Utilize WaitGroups para que seu programa não finalize antes de suas goroutines.
//     - Demonstre que há uma condição de corrida utilizando a flag -race
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