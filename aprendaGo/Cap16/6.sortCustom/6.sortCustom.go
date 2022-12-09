// - O sort que eu quero não existe. Quero fazer o meu.
// - Para isso podemos usar o func Sort do package sort. Vamos precisar de um sort.Interface.
//   - type Interface interface { Len() int; Less(i, j int) bool; Swap(i, j int) }
//
// - Ou seja, se tivermos um tipo que tenha esses métodos, ao executar sort.Sort(x) as funções que vão rodar são as minhas, não as funções pré-prontas como no exercício anterior.
// - E aí posso fazer do jeito que eu quiser.
// - Exemplo:
//   - struct carros: nome, consumo, potencia
//   - slice []carros{carro1, carro2, carro3} (Sort ordena *slices!*)
//   - tipo ordenarPorPotencia
//   - tipo ordenarPorConsumo
//   - tipo ordenarPorLucroProDonoDoPosto
package main

import (
	"fmt"
	"sort"
)

type carro struct{
    nome string
    consumo int
    potencia int
}

type ordenarPorPotencia []carro 
type ordenarPorConsumo []carro
type ordenarPorLucroProDonoDoPosto []carro


func (a ordenarPorPotencia) Len() int           { return len(a) }
func (a ordenarPorPotencia) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ordenarPorPotencia) Less(i, j int) bool { return a[i].potencia < a[j].potencia }

func (a ordenarPorConsumo) Len() int           { return len(a) }
func (a ordenarPorConsumo) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ordenarPorConsumo) Less(i, j int) bool { return a[i].consumo < a[j].consumo }

func (a ordenarPorLucroProDonoDoPosto) Len() int           { return len(a) }
func (a ordenarPorLucroProDonoDoPosto) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ordenarPorLucroProDonoDoPosto) Less(i, j int) bool { return a[i].consumo + a[i].potencia < a[j].consumo + a[j].potencia }

func main() {
    carro1 := carro{"Fusca", 10, 50}
    carro2 := carro{"Ferrari", 5, 100}
    carro3 := carro{"Gol", 15, 30}
    carro4 := carro{"Mclaren", 3, 101}
    carros := []carro{carro1, carro2, carro3, carro4}
    fmt.Println("Inicial:")
    fmt.Println("\t",carros)

    fmt.Println("Potencia:")
    sort.Sort(ordenarPorPotencia(carros))
    fmt.Println("\t",carros)
    
    fmt.Println("Consumo:")
    sort.Sort(ordenarPorConsumo(carros))
    fmt.Println("\t",carros)
    
    fmt.Println("Lucro para o posto(Potencia + Consumo):")
    sort.Sort(ordenarPorLucroProDonoDoPosto(carros))
    fmt.Println("\t",carros)

    
}