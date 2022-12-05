package main

import "fmt"

type tipo int

var y tipo = 42

func main() {
	x := 42
	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	x = int(y)
	fmt.Printf("%v %T\n", x, x)
}