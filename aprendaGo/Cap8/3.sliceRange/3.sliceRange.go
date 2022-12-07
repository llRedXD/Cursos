// - Slices:
//   - Tamanho: len(x)
//   - Índice específico: x[i] (0-based)
//
// - Para ver todos os itens de uma slice utilizamos o loop for com range.
// - Range significa alcance, faixa, extensão.
// - For range: for i, v := range x {}
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	for i, v := range slice {
		println("No índice ", i, " temos o valor: ", v)
	}
	
	fmt.Print("\n")

	// Caso queira Adicionar um item no slice
	slice = append(slice, 7)
	for i, v := range slice {
		println("No índice ", i, " temos o valor: ", v)
	}
	
	total := 4
	for _, v := range slice {
		total += v // Mesma coisa que total = total + valor
	}
	fmt.Println("Valor total é: ",total)

	

}