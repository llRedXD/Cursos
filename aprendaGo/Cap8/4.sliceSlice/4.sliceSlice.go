// - x[:], x[a:], x[:b], x[a:b]
// - "a" é incluso;
// - "b" não é.
// - Exemplo: cabeça magnética de um disco rígido (relógio, fita).
//   - Off-by-one error.
//
// - Go Playground: https://play.golang.org/p/i5ZOLKb3Fi
// - É fatiando que se deleta um item de uma slice. Na prática:
//   - x := append(x[:i], x[:i]...)
package main

import "fmt"

func main() {
	sabores := []string {"Abacaxi", "Calabresa", "Mussarela", "Dois Queijos"}

	fatia := sabores[0:2]
	fmt.Println(fatia)
	fatia2 := sabores[2:]
	fmt.Println(fatia2)

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}
	sabores = append(sabores[:2], sabores[3:]...) // Para deletarmos um slice criamos um slice selecionando todos os itens que queremos
	fmt.Println(sabores)

}