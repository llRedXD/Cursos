// - Crie um programa que lance 10 goroutines onde cada uma envia 10 números a um canal;
// - Tire estes números do canal e demonstre-os.
package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	f := 10

	go send(f, c)

	for v := range c {
		fmt.Println(v)
	}

}

func send(f int, c chan int) {
	var wg sync.WaitGroup
	for i := 0; i < f; i++ {
		wg.Add(1)
		go func() {
			total := 10
			for i := 1; i < total; i++ {
				time.Sleep(1e3)
				c <- i
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c)
}

// package main

// import "fmt"

// func main() {
// 	canal := make(chan int)
// 	for i := 0; i < 10; i++ {
// 		go func() {
// 			for j := 0; j < 10; j++ {
// 				canal <- j
// 			}
// 		}()
// 	}
// 	for k := 0; k < 100; k++ {
// 		fmt.Println(k, "\t", <-canal)
// 	}
// }
