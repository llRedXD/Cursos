// - Slices são feitas de arrays.
// - Elas são dinâmicas, podem mudar de tamanho.
// - Sempre que isso acontece, um novo array é criado e os dados são copiados.
// - É conveniente, mas tem um custo computacional.
// - Para otimizar as coisas, podemos utilizar make.
// - make([]T, len, cap)
// - "The length of a slice may be changed as long as it still fits within the limits of the underlying array; just assign it to a slice of itself. The capacity of a slice, accessible by the built-in function cap, reports the maximum length the slice may assume."
// - len(x), cap(x)
// - x[n] onde n é maior que len é out of range. Use append.
// - Append maior que cap modifica o array subjacente.
// - pkg/builtin/#append: "If it has sufficient capacity, the destination is resliced to accommodate the new elements. If it does not, a new underlying array will be allocated."
// - Effective Go.
package main

import "fmt"

func main() {
	slice := make([]int, 5, 10) // slice de inteiros, tamanho 5, capacidade 10

	slice[0], slice[1], slice[2], slice[3], slice[4] = 6,7,8,9,10

	slice = append(slice, 1,2,3,4,5,7) 
	fmt.Println(slice, len(slice), cap(slice))  // Toda vez que o slice cresce, o array subjacente cresce também no dobro do tamanho anterior. Mas somente se o slice estiver maior que a capacidade do array subjacente. Se o slice estiver menor que a capacidade do array subjacente, o array não cresce.
}