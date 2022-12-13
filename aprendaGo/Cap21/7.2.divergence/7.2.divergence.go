// - Divergência é o contrário de convergência :)
// - Na prática, exemplos:
//   - 1. Um stream vira centenas de go funcs que depois convergem.
//   - Dois canais.
//   - Uma func manda X números ao primeiro canal.
//   - Outra func faz um range deste canal, e para cada ítem lança uma go func que poe o retorno de trabalho() no canal dois.
//   - Trabalho() é um timer aleatório pra simular workload.
//   - Por fim, range canal dois demonstra os valores.
//   - 2. Com throttling! Ou seja, com um número máximo de go funcs.
//   - Ídem acima, mas a func que lança go funcs é assim:
//   - Cria X go funcs, cada uma com um range no primeiro canal que, para cada item, poe o retorno de trabalho() no canal dois.
package main

import (
	"fmt"
	// "math/rand"
	"sync"
	"time"
)

func main() {
    c1 := make(chan int)
    c2 := make(chan int)
	funcs := 5


    go manda(20, c1)
    go outra(funcs,c1,c2)

    for v := range c2 {
        fmt.Println(v)
    }

}

func manda(n int, c chan int)  {
    for i := 0; i < n; i++ {
        c <- i
    }
    close(c)
}

func outra(funcs int, c1,c2 chan int)  {
    var wg sync.WaitGroup

	for i := 0; i < funcs; i++ {
		wg.Add(1)
		go func() {
			for v := range c1{
				c2 <- trabalho(v)
			}
			wg.Done()
		}()
	}
    wg.Wait()
    close(c2)
}

func trabalho(n int) int {
    time.Sleep(time.Millisecond * 1000) //time.Duration(rand.Intn(1e3))
    return n
}