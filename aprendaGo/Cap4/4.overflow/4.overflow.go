package main

import "fmt"


func main() {
	x := uint16(65535)
	fmt.Println(x)
	// Se eu somar 1 a x, ele vai dar overflow e voltar a 0, pois o valor máximo de uint16 é 65535
	x++
	fmt.Println(x)
	// Se eu somar 2 a x, ele vai dar overflow e voltar a 1 e quando somar 3 vai voltar a 2 e assim por diante
	x++
	fmt.Println(x)
	// Mas se eu tentar declarar um valor maior que 65535, ele vai dar erro
	// var y uint16 = 65536
	// fmt.Println(y)
}