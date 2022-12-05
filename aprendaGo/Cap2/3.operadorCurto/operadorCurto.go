package main

import "fmt"

// Para declarar uma variável fora de uma função, temos que usar a palavra chave var
var x = 5

func main() {

	fmt.Printf("x : %v, %T\n\n", x, x) 
	// O operador curto de declaração só pode ser usado dentro de funções
	// O operador curto de declaração só pode ser usado caso as variáveis sejam novas, ou seja, não podem ser declaradas antes
	x := 10
	y := "hello"
	z := 10.2

	fmt.Printf("x : %v, %T\n", x, x) // %v é o valor da variável, %T é o tipo da variável
	fmt.Printf("y : %v, %T\n", y, y)
	fmt.Printf("z : %v, %T\n\n", z, z)
	
	// Caso eu queira modificar o valor de uma variável eu posso usar o operador curto de atribuição se eu tentar usar operador curto de declaração, o compilador vai dar um erro
	x = 20
	fmt.Printf("x : %v, %T\n\n", x, x)

	// Mas podemos usar o operador curto de declaração para declarar mais de uma variável assim não ira dar erro
	x, w := 30, 40
	fmt.Printf("x : %v, %T\n", x, x)
	fmt.Printf("w : %v, %T\n\n", w, w)

	// O nome das variáveis não podem começar com números, mas podem conter números, também não podem conter caracteres especiais, exceto o _ (underline) e não podem ser palavras reservadas(func, var, package, etc)

	// Podemos declarar expressões dentro do operador curto de declaração
	o := 10 + 10
	fmt.Printf("%o \n\n", o)

	// Operadores de comparação
	// == igual
	// != diferente
	// > maior
	// < menor
	// >= maior ou igual
	// <= menor ou igual
	q := 10 == 10
	fmt.Printf("%v\n\n", q)

}