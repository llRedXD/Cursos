// - Isso tudo aqui a gente já viu:
// - Toda slice tem um array subjacente.
// - Um slice é: um ponteiro/endereço para um array, mais len e cap (que é o len to array).
// - Exemplo:
//   - x := []int{...números}
//   - y := append(x[:i], x[:i]...)
//   - pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
//   - Ou seja, y utiliza o mesmo array subjacente que x.
//   - O que nos dá um resultado inesperado.
//
// - Ou seja, bom saber de antemão pra não ter que aprender na marra.
package main

import "fmt"

func main() {
	slice := []int {0,1,2,3,4,5}

	slice2 := append(slice[:2], slice[4:]...)

	fmt.Println(slice2) // Aqui temos o slice2 com os valores 0,1,4,5 e o slice com os valores 0,1,4,5,4,5
	
	fmt.Println(slice) // Se eu não quiser quebrar meu slice, eu posso fazer assim:

	slice3 := []int {0,1,2,3,4,5}

	slice3 = append(slice3[:2], slice3[4:]...)

	fmt.Println(slice3) // Aqui temos o slice3 com os valores 0,1,4,5 e o slice3 com os valores 0,1,4,5

}
