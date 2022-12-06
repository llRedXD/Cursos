// - For: inicialização, condição, pós
// - For: condição ("while")
// - For: ...ever? (http servers)
// - For: break
// - golang.org/ref/spec#For_statements, Effective Go
// - (Range vem mais pra frente.)
package main

func main() {
	for y:= 0;y < 10; y++ { // for padrão (inicialização, condição, pós)
		println(y)
	}
	x := 0
	for x < 10 { // for tipo while
		println(x)
		x++
	}
	// for { // for infinito
	// 	println("loop infinito") 
	// }
	// Break serve para quebrar um loop
}