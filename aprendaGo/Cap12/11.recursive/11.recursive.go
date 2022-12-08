// - WP: "A aplicação mais comum de recursão é em matemática e ciência da computação, onde uma função que está sendo definida é aplicada dentro de sua própria definição."
// - Exemplos de recursividade: Fractais, matrioscas, efeito Droste (o efeito produzido por uma imagem que aparece dentro dela própria), GNU (“GNU is Not Unix”), etc.
// - No estudo de funções: é uma função que chama a ela própria.
// - Exemplo: fatoriais.
//     - 4! = 4 * 3 * 2 * 1 (e no zero, deu.)
//     - Com recursividade. Go Playground: https://play.golang.org/p/ujsLnUhRp_
//     - Com loops. Go Playground: https://play.golang.org/p/F2VsUjYVhc
package main

import "fmt"

func main() {
	fmt.Println(fatorial(6))
}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}