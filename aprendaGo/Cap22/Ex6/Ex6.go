// - Escreva um programa que coloque 100 números em um canal, retire os números do canal, e demonstre-os.
package main

import "fmt"

func main() {
	c := make(chan int)

	go send(c)

	for v := range c {
		fmt.Println(v)
	}

}

func send(c chan int) {
	total := 101
	for i := 1; i < total; i++ {
		c <- i
	}
	close(c)
}
