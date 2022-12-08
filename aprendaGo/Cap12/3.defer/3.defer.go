// - Funções são ótimas pois tornam nosso código modular. Podemos alterar partes do nosso programa sem afetar o resto!
// - Uma declaração defer chama uma função cuja execução ocorrerá no momento em que a função da qual ela faz parte finalizar.
// - Essa finalização pode ocorrer devido a um return, ao fim do code block da função, ou no caso de pânico em uma goroutine correspondente.
// - "Deixa pra última hora!"
// - ref/spec
// - Sempre usamos para fechar um arquivo após abri-lo.
package main

import "fmt"

func main() {
	defer fmt.Println("Antes") // A função fmt.Println não será executada agora, mas sim quando a função main terminar.
	defer fmt.Println("Um pouco antes")
	fmt.Println("Depois")
	fmt.Println("Um pouco depois")
}