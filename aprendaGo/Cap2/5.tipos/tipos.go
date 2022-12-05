package main

import "fmt"

// Os tipos podem ser declarados de duas formas, de forma a ser deduzido pelo compilador ou de forma explícita
var x = 10 // forma implícita
var y string = "string" // forma explícita

// Caso eu declare somente o tipo o valor so atribuído dentro de um escopo, como uma função senão for declarado o valor o valor padrão é o zero 
var z int 

func main() {

	fmt.Println(x, y, z)
}

//  Tipos Compostos
// - Array
// - Slice
// - Struct
// - Map
// - Pointer
// - Function
// - Interface

// Esse tipo de dado é composto por um composto de tipos primitivos e criado pelo usuário
// O ato de definir, cria, estruturar tipos compostos é chamado de composição