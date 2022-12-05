// - Strings são sequencias de bytes.
// - Imutáveis.
// - Uma string é um "slice of bytes" (ou, em português, uma fatia de bytes).
// - Na prática:
//     - %v %T
//     - Raw string literals
//     - Conversão para slice of bytes: []byte(x)
//     - %#U, %#x

package main

import "fmt"

func main() {
	s := "Olá, mundo!"
	fmt.Printf("%v\n%T\n", s, s)
	// Caso queira imprimir uma string literal, use raw string literals.

	t := `Olá, 
	

	mundo!`
	fmt.Printf("%v\n%T\n", t, t)
	// Para converter uma string para um slice of bytes, use []byte(x).
	
	sb := []byte(s)
	fmt.Printf("%v\n%T\n", sb, sb)
	
	// Com o slice of bytes, é possível acessar cada byte da string. Assim podemos criar um loop para imprimir cada byte.
	for _, v := range s {
		// %#U imprime o valor em decimal, %#x imprime o valor em hexadecimal.
		fmt.Printf("%v - %T - %#U - %#x\n", v, v, v, v)
	}

	fmt.Println("")

	for i := 0; i < len(s); i++ {
		fmt.Printf("%v - %T - %#U - %#x\n", s[i], s[i], s[i], s[i])
	}
}