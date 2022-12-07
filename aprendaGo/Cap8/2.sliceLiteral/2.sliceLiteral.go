// - O que são tipos de dados compostos?
//   - Wikipedia: Composite_data_type
//   - Effective Go: Composite literals
//   - ref/spec: Composite literals
//
// - Uma slice agrupa valores de um único tipo.
// - Criando uma slice: literal composta → x := []type{values}
package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(slice)
	
	slice2 := append(slice, 7)
	fmt.Println(slice2)
	slice2[2] = 202020
	fmt.Println(slice2)
	// slice2[6] = 202020 Não funciona, pois o slice2 tem tamanho 7

}