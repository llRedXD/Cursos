// - Sort serve para ordenar slices.
//     - Como faz?
//     - golang.org/pkg/ → sort
//     - godoc.org/sort → examples
//     - Sort altera o valor original!
// - Exemplo: Ints, Strings.
package main

import (
	"fmt"
	"sort"
)

func main() {
    s := []string{"Zeno", "John", "Al", "Jenny"}
    fmt.Println(s)
    sort.Strings(s)
    fmt.Println(s)
    i := []int{1,5,48,69,7,9}
    fmt.Println(i)
    sort.Ints(i)
    fmt.Println(i)

}