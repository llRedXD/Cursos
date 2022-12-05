package main

import "fmt"

type tipo int // tipos criados não interagem com tipos nativos

var x tipo

func main() {
	fmt.Printf("%T\n", x)
	y := 42
	fmt.Printf("%T\n", y)
}